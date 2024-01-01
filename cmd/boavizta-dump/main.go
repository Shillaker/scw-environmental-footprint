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
	"id", // Type of instance e.g. c1.medium
	"manufacturer", // Just Scaleway
	"CASE.case_type", // Always rack
	"year", // Year of release?
	"vcpu", // Number of vCPUs the instance has, e.g. 2
	"platforme_vcpu", // Number of threads the underlying host has e.g. 48
	"CPU.units", // Number of CPUs installed on the machine e.g. 2
	"CPU.core_units", // Number of cores the CPU has e.g. 12
	"CPU.name", // CPU model e.g. Xeon E5-2651 v2
	"CPU.manufacturer", // CPU manufacturer e.g. intel
	"CPU.model_range", // CPU range e.g. xeon-e5
	"CPU.family", // CPU family e.g. ivybridge
	"CPU.tdp", // CPU TDP e.g. 95
	"CPU.manufacture_date", // CPU date built e.g. 2008
	"instance.ram_capacity", // Ram available to VM (GB) e.g. 2
	"RAM.capacity", // Ram available to host e.g. 2
	"RAM.units", // Ram units installed in host e.g. 21
	"SSD.units", // SSDs installed on host e.g. 1
	"SSD.capacity", // SSD capacity e.g. 350
	"HDD.units", // HDDs installed on host e.g. 1
	"HDD.capacity", // HDD capacity (GB) e.g. 350
	"GPU.name", // Name of GPU available e.g. T4
	"GPU.units", // Number of GPUs units available e.g. 8
	"GPU.TDP", // GPU TDP e.g. 70
	"GPU.memory_capacity", // GPU memory (GB) e.g. 16
	"POWER_SUPPLY.units", // Number of PSUs on host with semicolons e.g. 2;2;2
	"POWER_SUPPLY.unit_weight", // Weighting of each PSU e.g. 2.99;1;5
	"USAGE.instance_per_server", // Number of VMs per host e.g. 24
	"USAGE.time_workload", // Workload? e.g. 50;0;100
	"USAGE.use_time_ratio", // Usage time ratio? e.g. 1
	"USAGE.hours_life_time", // Hours lifetime e.g. 35040
	"USAGE.other_consumption_ratio", // Consumption ratio? e.g. 0.33;0.2;0.6
	"USAGE.overcommited", // Whether host is overcommitted e.g. 0
	"Warnings", // Notes e.g. RAM.capacity not verified
}


type BoaviztaCsvLine struct {
	id                  string
	manufacturer        string

	instance			model.InstanceBaseServer
}

func (b *BoaviztaCsvLine) toRow() []string {
	return []string{
		b.id,
		scalewayManufacturer,
		caseType,
		fmt.Sprintf("%v", defaultYear),
		fmt.Sprintf("%v", b.instance.VCpus),
		fmt.Sprintf("%v", b.instance.Server.Cpus
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

		err = writer.Write(bvLine.toRow())
		if (err != nil) {
			fmt.Errorf("could not write CSV line %v", err)
		}
	}

}
