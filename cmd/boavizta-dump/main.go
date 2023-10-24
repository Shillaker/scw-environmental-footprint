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
	outFile              = "boavizta.csv"
	scalewayManufacturer = "SCW"
	caseType             = "rack"
	defaultYear          = 2022
)

var (
	outPath = filepath.Join(outDir, outFile)
)

var bvHeaders = []string{
	"id",
	"manufacturer",
	"CASE.case_type",
	"year",
	"vcpu",
	"platforme_vcpu",
	"CPU.units",
	"CPU.core_units",
	"CPU.name",
	"CPU.manufacturer",
	"CPU.model_range",
	"CPU.family",
	"CPU.tdp",
	"CPU.manufacture_date",
	"instance.ram_capacity",
	"RAM.capacity",
	"RAM.units",
	"SSD.units",
	"SSD.capacity",
	"HDD.units",
	"HDD.capacity",
	"GPU.name",
	"GPU.units",
	"GPU.TDP",
	"GPU.memory_capacity",
	"POWER_SUPPLY.units",
	"POWER_SUPPLY.unit_weight",
	"USAGE.instance_per_server",
	"USAGE.time_workload",
	"USAGE.use_time_ratio",
	"USAGE.hours_life_time",
	"USAGE.other_consumption_ratio",
	"USAGE.overcommited",
	"Warnings",
}


type BoaviztaCsvLine struct {
	id                  string
	manufacturer        string
	caseType            string
	year                int32
	vmVCpu              int32
	hostVCpu            int32
	cpuCount            int32
	cpuCores            int32
	cpuName             string
	cpuManufacturer     string
	cpuModelRange       string
	cpuFamily           string
	cpuTdp              int32
	cpuManufactureDate  string
	instanceRamCapacity string
	ramCapacity         int32
	ramCount            int32
	ssdCount            int32
	ssdCapacity         int32
	hddCount            int32
	hddCapacity         int32
	gpuName             string
	gpuCount            int32
	gpuTdp              int32
	gpuMemoryCapacity   int32
	psuCount            int32
}

func (b *BoaviztaCsvLine) toRow() []string {
	return []string{
		b.id,
		b.manufacturer,
		b.caseType,
		fmt.Sprintf("%v", b.year),
		fmt.Sprintf("%v", b.vmVCpu),
		fmt.Sprintf("%v", b.hostVCpu),
		fmt.Sprintf("%v", b.cpuCount),
		fmt.Sprintf("%v", b.cpuCores),
		b.cpuName,
		b.cpuManufacturer,
		b.cpuModelRange,
		b.cpuFamily,
		fmt.Sprintf("%v", b.cpuTdp),
		b.cpuManufactureDate,
		b.instanceRamCapacity,
		fmt.Sprintf("%v", b.ramCapacity),
		fmt.Sprintf("%v", b.ramCount),
		fmt.Sprintf("%v", b.ssdCount),
		fmt.Sprintf("%v", b.ssdCapacity),
		fmt.Sprintf("%v", b.hddCount),
		fmt.Sprintf("%v", b.hddCapacity),
		b.gpuName,
		fmt.Sprintf("%v", b.gpuCount),
		fmt.Sprintf("%v", b.gpuTdp),
		fmt.Sprintf("%v", b.gpuMemoryCapacity),
		fmt.Sprintf("%v", b.psuCount),
	}
}

func main() {
	util.InitLogging()
	err := util.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	// Create the directory
	os.Mkdir(outDir, os.FileMode(0775))

	// Create the file with headings
	f, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("failed to open file at %v: %v", outPath, err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	writer.Write(bvHeaders)

	// Print out Boavizta config
	for name, instance := range model.InstanceServerMapping {
		bvLine := BoaviztaCsvLine{
			id:           name,
			manufacturer: scalewayManufacturer,
			caseType:     caseType,
			year:         defaultYear,
			vmVCpu:       instance.VCpus,
			cpuCount:     1,
			psuCount:     1,
			ssdCount:     1,
		}

		writer.Write(bvLine.toRow())
	}
}
