package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"gitlab.infra.online.net/paas/carbon/model"
	"gitlab.infra.online.net/paas/carbon/util"
)

const (
	outDir               = "output"
	instancesOutFile     = "instances.csv"
	serversOutFile       = "servers.csv"
	scalewayManufacturer = "SCW"
	caseType             = "rack"
	defaultYear          = 2022
)

var (
	instancesOutPath = filepath.Join(outDir, instancesOutFile)
	serversOutPath   = filepath.Join(outDir, serversOutFile)
)

var instancesHeaders = []string{
	"id",          // Identifier
	"vcpu",        // Number of vCPUs
	"memory",      // GiB of memory
	"ssd_storage", // GiB of SSD storage
	"hdd_storage", // GiB of HDD storage
	"gpu_units",   // Number of GPUs
	"platform",    // Underlying platform/bare metal server
	"source",      // Link to information
}

var serversHeaders = []string{
	"id",
	"manufacturer",
	"CASE.case_type",
	"CPU.units",
	"CPU.core_units",
	"CPU.model_range",
	"CPU.die_size_per_core",
	"CPU.name",
	"CPU.threads",
	"CPU.manufacturer",
	"CPU.tdp",
	"CPU.family",
	"RAM.units",
	"RAM.density",
	"RAM.capacity",
	"SSD.units",
	"SSD.capacity",
	"SSD.density",
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

	for _, server := range model.InstanceBaseServers {
		row := []string{
			server.Name, // ID
			"",          // Manufacturer
			"rack",
			fmt.Sprintf("%v", server.TotalCpuUnits()), // CPU units
			fmt.Sprintf("%v", server.TotalCores()),    // CPU cores
			server.Cpus[0].Model,                      // CPU model
			"",                                        // CPU die size
			server.Cpus[0].Name,                       // CPU name
			fmt.Sprintf("%v", server.TotalCores()),    // CPU threads - TODO: hyper threading?
			server.Cpus[0].Manufacturer,               // CPU manufacturer
			"",                                        // CPU TDP
			server.Cpus[0].Family,                     // CPU.family,
			fmt.Sprintf("%v", server.TotalRamUnits()), // RAM units
			"", // RAM density
			fmt.Sprintf("%v", server.TotalRamCapacity()/1024), // RAM capacity (GiB)
			fmt.Sprintf("%v", server.TotalSsdUnits()),         // SSD units
			fmt.Sprintf("%v", server.TotalSsdCapacity()),      // SSD capacity
			"", // SSD density
			fmt.Sprintf("%v", server.TotalHddUnits()),    // HDD units
			fmt.Sprintf("%v", server.TotalHddCapacity()), // HDD capacity
			"0",                         // GPU units
			"",                          // GPU name
			"0",                         // GPU memory_capacity
			"1",                         // POWER_SUPPLY units
			"",                          // POWER_SUPPLY unit_weight
			"50;0;100",                  // USAGE time_workload - TODO: what is this?
			"1",                         // USAGE use_time_ratio - TODO: what is this?
			fmt.Sprintf("%v", 24*365*8), // USAGE hours_life_time - TODO: pick suitable lifespan
			"0.33;0.2;0.6",              // USAGE other_consumption_ratio - TODO: what is this?
			"",                          // WARNINGS
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
