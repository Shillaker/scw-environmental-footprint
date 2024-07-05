package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/shillaker/scw-environmental-footprint/model"
	"github.com/shillaker/scw-environmental-footprint/util"
)

// See servers doc for info on fields and format:
// https://github.com/Boavizta/boaviztapi/blob/main/docs/docs/contributing/server.md
const (
	dataDir                      = "data"
	instancesOutFile             = "instances.csv"
	serversOutFile               = "servers.csv"
	caseType                     = "rack"
	lifespanHours                = model.DefaultLifespanYears * 365 * 24
	defaultTimeWorkload          = "50;0;100"
	defaultUseTimeRatio          = "1"
	defaultOtherConsumptionRatio = "0.33;0.2;0.6"
)

var (
	outDir           = filepath.Join(dataDir, "output")
	instancesOutPath = filepath.Join(outDir, instancesOutFile)
	serversOutPath   = filepath.Join(outDir, serversOutFile)
)

// Instance headers as listed here:
// https://github.com/Boavizta/boaviztapi/blob/main/docs/docs/contributing/cloud_instance.md
var instancesHeaders = []string{
	"id",
	"vcpu",
	"memory",
	"ssd_storage",
	"hdd_storage",
	"gpu_units",
	"platform",
	"source",
}

// Server headers as listed here:
// https://github.com/Boavizta/boaviztapi/blob/main/docs/docs/contributing/server.md
var serversHeaders = []string{
	"id",
	"manufacturer",
	"CASE.case_type",
	"CPU.units",
	"CPU.core_units",
	"CPU.die_size_per_core",
	"CPU.name",
	"CPU.vcpu",
	"RAM.units",
	"RAM.capacity",
	"SSD.units",
	"SSD.capacity",
	"HDD.units",
	"HDD.capacity",
	"GPU.units",
	"GPU.name",
	"GPU.memory_capacity",
	"POWER_SUPPLY.units",
	"POWER_SUPPLY.unit_weight",
	"USAGE.time_workload",
	"USAGE.use_time_ratio",
	"USAGE.hours_life_time",
	"USAGE.other_consumption_ratio",
	"WARNINGS",
}

func writeInstances() error {
	instancesFile, err := os.Create(instancesOutPath)
	if err != nil {
		log.Errorf("failed to open instances file at %v: %v", instancesOutPath, err)
		return err
	}
	defer instancesFile.Close()

	// Create the servers file with headings
	instancesWriter := csv.NewWriter(instancesFile)
	instancesWriter.Write(instancesHeaders)
	defer instancesWriter.Flush()

	for name, instance := range model.InstanceServerMapping {
		row := []string{
			name,
			fmt.Sprintf("%v", instance.VCpus),
			fmt.Sprintf("%v", instance.RamGiB),
			fmt.Sprintf("%v", instance.SsdGiB),
			fmt.Sprintf("%v", instance.HddGiB),
			fmt.Sprintf("%v", instance.Gpus),
			fmt.Sprintf("%v", instance.Server.Name),
			"",
		}

		err = instancesWriter.Write(row)
		if err != nil {
			log.Errorf("ERROR: could not write instance CSV line")
			return err
		}
	}

	return nil
}

func writeServers() error {
	serversFile, err := os.Create(serversOutPath)
	if err != nil {
		log.Errorf("failed to open servers file at %v: %v", serversOutPath, err)
		return err
	}
	defer serversFile.Close()

	serversWriter := csv.NewWriter(serversFile)
	serversWriter.Write(serversHeaders)
	defer serversWriter.Flush()

	for _, server := range model.VirtualMachines {
		row := []string{
			server.Name, // ID
			"",          // Manufacturer
			caseType,
			fmt.Sprintf("%v", server.TotalCpuUnits()),         // CPU units
			fmt.Sprintf("%v", server.CpuCores()),              // Cores per CPU
			"",                                                // CPU die size per core
			server.CpuName(),                                  // CPU name
			fmt.Sprintf("%v", server.VCpuPerCpu),              // Number of vCPUs per CPU
			fmt.Sprintf("%v", server.TotalRamUnits()),         // RAM units
			fmt.Sprintf("%v", server.TotalRamCapacity()/1024), // RAM capacity (GiB)
			fmt.Sprintf("%v", server.TotalSsdUnits()),         // SSD units
			fmt.Sprintf("%v", server.TotalSsdCapacity()/1024), // SSD capacity
			fmt.Sprintf("%v", server.TotalHddUnits()),         // HDD units
			fmt.Sprintf("%v", server.TotalHddCapacity()/1024), // HDD capacity
			fmt.Sprintf("%v", server.TotalGpuUnits()),         // GPU units
			server.GpuName(),                                  // GPU name
			fmt.Sprintf("%v", server.TotalGpuMemory()/1024),   // GPU memory_capacity
			fmt.Sprintf("%v", server.PowerSupply.Units),       // POWER_SUPPLY units
			fmt.Sprintf("%v", server.PowerSupply.WeightKg),    // POWER_SUPPLY unit_weight
			defaultTimeWorkload,                               // USAGE time_workload
			"1",                                               // USAGE use_time_ratio
			fmt.Sprintf("%v", lifespanHours),                  // USAGE hours_life_time
			defaultOtherConsumptionRatio,                      // USAGE other_consumption_ratio
			"",                                                // WARNINGS
		}

		err = serversWriter.Write(row)
		if err != nil {
			log.Errorf("ERROR: could not write server CSV line")
			return err
		}
	}

	return nil
}

func main() {
	util.InitLogging()
	err := util.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	os.Mkdir(outDir, os.FileMode(0775))

	err = writeInstances()
	if err != nil {
		log.WithError(err).Fatalf("Failed writing instances")
	}

	err = writeServers()
	if err != nil {
		log.WithError(err).Fatalf("Failed writing servers")
	}
}
