package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/shillaker/scw-environmental-footprint/model"
	"github.com/shillaker/scw-environmental-footprint/util"
)

const (
	dataDir          = "data"
	instancesOutFile = "instances.csv"
	serversOutFile   = "servers.csv"
)

var instancesHeaders = []string{
	"id",
	"vcpu",
	"ram_gib",
	"ssd_gib",
	"hdd_gib",
	"gpu_units",
	"base_server",
}

var CpuHeaders = []string {
	"id",
	"manufacturer",
	"model",
	"freq_hz",
	"cores",
	"threads",
}

var baseServerHeaders = []string {
	"id",
	"product_category",
	"manufacturer",
	"model",
	"cpu_id",
	"ram_dimm_count",
	"ram_dimm_size",
	"ram_type",
	"ssd_size_gib",
	"ssd_count",
	"hdd_size_gib",
	"psu_count",

}

var (
	outDir           = filepath.Join(dataDir, "output")
	instancesOutPath = filepath.Join(outDir, instancesOutFile)
	serversOutPath   = filepath.Join(outDir, serversOutFile)
)

func writeInstances() error {
	yamlData, err := yaml.Marshal(model.InstanceServerMapping)
	if err != nil {
		log.Errorf("failed to serialise instances to YAML %v", err)
		return err
	}

	instancesFile, err := os.Create(instancesOutPath)
	if err != nil {
		log.Errorf("failed to open instances file at %v: %v", instancesOutPath, err)
		return err
	}
	instancesFile.Close()

	err = ioutil.WriteFile(instancesOutPath, yamlData, 0)
	if err != nil {
		log.Errorf("failed to write instances at %v: %v", instancesOutPath, err)
		return err
	}

	return nil
}

func writeServers() error {
	yamlData, err := yaml.Marshal(model.VirtualMachines)
	if err != nil {
		log.Errorf("failed to serialise servers to YAML %v", err)
		return err
	}

	serversFile, err := os.Create(serversOutPath)
	if err != nil {
		log.Errorf("failed to open servers file at %v: %v", serversOutPath, err)
		return err
	}
	defer serversFile.Close()

	err = ioutil.WriteFile(serversOutPath, yamlData, 0)
	if err != nil {
		log.Errorf("failed to write servers at %v: %v", serversOutPath, err)
		return err
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
