package boavizta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/shillaker/scw-environmental-footprint/model"
	"github.com/shillaker/scw-environmental-footprint/util"
	"github.com/spf13/viper"
)

type BoaviztaImpactCalculator struct {
	Host string
	Port string
}

func mapRegionToBoaviztaRegion(regionIn string) (string, error) {
	switch regionIn {
	case model.RegionFrance:
		return BoaviztaRegionFrance, nil
	case model.RegionNetherlands:
		return BoaviztaRegionNetherlands, nil
	case model.RegionPoland:
		return BoaviztaRegionPoland, nil
	default:
		return "", fmt.Errorf("Unsupported region for Boavizta: %v", regionIn)
	}
}

func mapServerUsageToBoaviztaModel(serverUsage model.ServerUsage) (BoaviztaServerRequest, error) {
	request := BoaviztaServerRequest{}

	// We assume all CPUs are the same
	if len(serverUsage.Server.Cpus) > 1 {
		return request, fmt.Errorf("Multiple CPU types not supported with Boavizta")
	}

	cpu := serverUsage.Server.Cpus[0]
	request.Configuration.Cpu = BoaviztaCpu{
		Family:    cpu.Name,
		Units:     int32(cpu.Units),
		CoreUnits: int32(cpu.CoreUnits),
		Tdp:       int32(cpu.Tdp),
	}

	// RAM
	for _, ram := range serverUsage.Server.Rams {
		request.Configuration.Ram = append(request.Configuration.Ram, BoaviztaRam{
			CapacityGib:  int32(ram.CapacityMib / 1024),
			Units:        1,
			Manufacturer: ram.Manufacturer,
		})
	}

	// SSD
	for _, ssd := range serverUsage.Server.Ssds {
		request.Configuration.Disk = append(request.Configuration.Disk, BoaviztaDisk{
			Type:         "ssd",
			CapacityGib:  int32(ssd.CapacityMib / 1024),
			Units:        int32(ssd.Units),
			Manufacturer: ssd.Manufacturer,
		})
	}

	// HDD
	for _, hdd := range serverUsage.Server.Hdds {
		request.Configuration.Disk = append(request.Configuration.Disk, BoaviztaDisk{
			Type:        "hdd",
			CapacityGib: int32(hdd.CapacityMib / 1024),
			Units:       int32(hdd.Units),
		})
	}

	// PSU
	request.Configuration.PowerSupply = BoaviztaPowerSupply{
		Units: int32(serverUsage.Server.PowerSupply.Units),
	}

	// Always rack
	request.Model.Type = "rack"

	// Here we specify a consumption profile, which can be used to estimate the energy consumption from the consumption profile for the given devices.
	// For example, we may say, for 100% of the specified time, the server was running at 50% load
	request.Usage.TimeWorkload = []BoaviztaTimeWorkload{
		{
			LoadPercentage: serverUsage.Usage.LoadPercentage,
			TimePercentage: 100,
		},
	}

	// Note need to round up here because request takes an integer
	request.Usage.HoursUseTime = serverUsage.Usage.TimeHoursRoundedUp()

	var err error
	request.Usage.UsageLocation, err = mapRegionToBoaviztaRegion(serverUsage.Usage.Region)

	return request, err
}

func mapBoaviztaResponseToImpact(response BoaviztaServerResponse, hostShare float32) model.ImpactServerUsage {
	impact := model.ImpactServerUsage{}
	impact.Impacts = make(map[string]model.Impact, 3)

	impact.Impacts["adp"] = model.Impact{
		Manufacture: response.Impacts.AdpImpact.Manufacture / hostShare,
		Use:         response.Impacts.AdpImpact.Use / hostShare,
		Unit:        response.Impacts.AdpImpact.Unit,
	}

	impact.Impacts["gwp"] = model.Impact{
		Manufacture: response.Impacts.GwpImpact.Manufacture / hostShare,
		Use:         response.Impacts.GwpImpact.Use / hostShare,
		Unit:        response.Impacts.GwpImpact.Unit,
	}
	impact.Impacts["pe"] = model.Impact{
		Manufacture: response.Impacts.PeImpact.Manufacture / hostShare,
		Use:         response.Impacts.PeImpact.Use / hostShare,
		Unit:        response.Impacts.PeImpact.Unit,
	}

	return impact
}

func (b *BoaviztaImpactCalculator) getBoaviztaUrl() string {
	return fmt.Sprintf("http://%v:%v", b.Host, b.Port)
}

func (b *BoaviztaImpactCalculator) getBoaviztaServerUrl() string {
	bvUrl := b.getBoaviztaUrl()

	// Note need for trailing slash here
	return fmt.Sprintf("%v/v1/server/?verbose=true&allocation=%v", bvUrl, BoaviztaAllocationTypeLinear)
}

func (b *BoaviztaImpactCalculator) CalculateServerImpact(serverUsage []model.ServerUsage) (model.ImpactServerUsage, error) {
	var impacts []model.ImpactServerUsage
	var empty model.ImpactServerUsage

	for _, s := range serverUsage {
		if s.HostShare == 0 {
			return empty, fmt.Errorf("Must set server host share")
		}

		request, err := mapServerUsageToBoaviztaModel(s)
		if err != nil {
			return empty, err
		}

		url := b.getBoaviztaServerUrl()

		log.Debugf("Making Boavizta request to %s", url)

		impact := model.ImpactServerUsage{}

		requestJson, err := json.Marshal(request)

		if err != nil {
			return empty, err
		}

		// Make the request to Boavizta
		responseBody := bytes.NewBuffer(requestJson)
		resp, err := http.Post(url, "application/json", responseBody)

		if err != nil {
			return empty, err
		}

		defer resp.Body.Close()

		// Read the response
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return empty, err
		}

		if resp.StatusCode != 200 {
			log.Errorf("Error making Boavizta request (%v): %v", resp.StatusCode, body)
			return empty, fmt.Errorf("Boavizta request failed")
		}

		// Map resonse to model
		var response BoaviztaServerResponse
		json.Unmarshal(body, &response)

		impact = mapBoaviztaResponseToImpact(response, s.HostShare)
		impacts = append(impacts, impact)
	}

	result := model.CombineImpacts(impacts)

	// Add equivalents
	result.EquivalentsManufacture = model.CalculateEquivalentCO2E(result.Impacts["gwp"].Manufacture)
	result.EquivalentsUse = model.CalculateEquivalentCO2E(result.Impacts["gwp"].Use)

	return result, nil
}

func NewBoaviztaImpactCalculator() (*BoaviztaImpactCalculator, error) {
	util.InitConfig()
	calc := &BoaviztaImpactCalculator{
		Host: viper.GetString("boavizta.host"),
		Port: viper.GetString("boavizta.port"),
	}

	log.Infof("Initialised Boavizta at %s:%s", calc.Host, calc.Port)

	return calc, nil
}
