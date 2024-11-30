package resilio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/shillaker/scw-environmental-footprint/model"
	"github.com/shillaker/scw-environmental-footprint/util"
	"github.com/spf13/viper"
)

type ResilioImpactCalculator struct {
	BaseUrl string
	Token   string
}

func mapRegionToResilioRegion(regionIn string) (string, error) {
	switch regionIn {
	case model.RegionFrance:
		return ResilioRegionFrance, nil
	case model.RegionNetherlands:
		return ResilioRegionNetherlands, nil
	case model.RegionPoland:
		return ResilioRegionPoland, nil
	default:
		return "", fmt.Errorf("Unsupported region for Resilio: %v", regionIn)
	}
}

func mapServerUsageToResilioModel(serverUsage model.ServerUsage) (ResilioRackServerUsage, error) {
	request := ResilioRackServerUsage{}

	request.UsagePercent = uint32(serverUsage.Usage.LoadPercentage)
	request.RackUnit = 1

	resilioRegion, err := mapRegionToResilioRegion(serverUsage.Usage.Region)
	if err != nil {
		return request, err
	}

	request.Usage = ResilioPowerUsage{
		DeltaTHour: uint32(serverUsage.Usage.TimeHoursRoundedUp()),
		PowerWatt:  serverUsage.Server.PowerSupply.Watts,
		Geography:  resilioRegion,
		// YearlyElectricityConsumption: serverUsage.Server.PowerSupply.YearlyConsumptionKwh(),
	}

	// Set all fields, Resilio API doesn't like null values
	request.Cpus = make([]ResilioCpu, 0)
	request.Gpus = make([]ResilioGpu, 0)
	request.Psus = make([]ResilioPsu, 0)
	request.Rams = make([]ResilioRam, 0)
	request.Ssds = make([]ResilioSsd, 0)

	for _, cpu := range serverUsage.Server.Cpus {
		request.Cpus = append(request.Cpus, ResilioCpu{
			Name: cpu.Name,
		})
	}

	for _, gpu := range serverUsage.Server.Gpus {
		request.Gpus = append(request.Gpus, ResilioGpu{
			Name: gpu.Name,
		})
	}

	for _, ram := range serverUsage.Server.Rams {
		request.Rams = append(request.Rams, ResilioRam{
			SizeGb: ram.CapacityMib / 1024,
		})
	}

	request.Hdds.Quantity = uint32(len(serverUsage.Server.Hdds))

	for _, ssd := range serverUsage.Server.Ssds {
		request.Ssds = append(request.Ssds, ResilioSsd{
			SizeGb:     ssd.CapacityMib / 1024,
			Technology: ssd.Technology,
			Casing:     ssd.Casing,
		})
	}

	request.Psus = append(request.Psus, ResilioPsu{
		PowerWatt: serverUsage.Server.PowerSupply.Watts,
	})

	return request, nil
}

func mapResilioResponseToImpact(response ResilioServerResponse) model.ImpactServerUsage {
	impact := model.ImpactServerUsage{}

	impact.Impacts = make(map[string]model.Impact)

	impact.Impacts["gwp"] = model.Impact{
		Manufacture: response.Results.Inventory.PerLcStep.GWPEmbedded(),
		Use:         response.Results.Inventory.PerLcStep.GWPUse(),
		Unit:        "kgCO2e",
	}

	impact.Impacts["adp"] = model.Impact{
		Manufacture: response.Results.Inventory.PerLcStep.ADPEmbedded(),
		Use:         response.Results.Inventory.PerLcStep.ADPUse(),
		Unit:        "kgSbeq",
	}

	impact.Impacts["pe"] = model.Impact{
		Manufacture: response.Results.Inventory.PerLcStep.PEEmbedded(),
		Use:         response.Results.Inventory.PerLcStep.PEUse(),
		Unit:        "MJ",
	}

	return impact
}

func (b *ResilioImpactCalculator) CalculateServerImpact(serverUsage []model.ServerUsage) (model.ImpactServerUsage, error) {
	var empty model.ImpactServerUsage

	request := ResilioRackServerRequest{}
	for _, s := range serverUsage {
		resilioUsage, err := mapServerUsageToResilioModel(s)
		if err != nil {
			return empty, err
		}

		request.ServerData = append(request.ServerData, resilioUsage)
	}

	url := fmt.Sprintf("%v/rack_server", b.BaseUrl)
	log.Infof("Making Resilio request to %s", url)

	requestJson, err := json.Marshal(request)

	if err != nil {
		return empty, err
	}

	reqBody := bytes.NewBuffer(requestJson)
	log.Debugf("Resilio request body:\n %s\n", reqBody.String())

	// Make the request
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return empty, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", b.Token)

	resp, err := client.Do(req)

	if err != nil {
		return empty, err
	}

	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return empty, err
	}

	if resp.StatusCode != 200 {
		bs := string(body)
		log.Errorf("Error making Resilio request (%v): %v", resp.StatusCode, bs)
		return empty, fmt.Errorf("Resilio request failed")
	}

	// Map response to model
	var response ResilioServerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return empty, err
	}

	result := mapResilioResponseToImpact(response)

	// Add equivalents
	result.EquivalentsManufacture = model.CalculateEquivalentCO2E(result.Impacts["gwp"].Manufacture)
	result.EquivalentsUse = model.CalculateEquivalentCO2E(result.Impacts["gwp"].Use)

	return result, nil
}

func NewResilioImpactCalculator() (*ResilioImpactCalculator, error) {
	err := util.InitConfig()
	if err != nil {
		return nil, err
	}

	calc := &ResilioImpactCalculator{
		BaseUrl: viper.GetString("resilio.base_url"),
		Token:   viper.GetString("resilio.token"),
	}

	log.Infof("Initialised Resilio at %s with token %s", calc.BaseUrl, calc.Token)

	return calc, nil
}
