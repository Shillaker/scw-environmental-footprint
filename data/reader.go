package data

import (
	_ "embed"

	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"github.com/shillaker/scw-environmental-footprint/model"
)

//go:embed source/elastic_metal.yaml
var elasticMetalData []byte

//go:embed source/apple_silicon.yaml
var appleSiliconData []byte

//go:embed source/instances.yaml
var instancesData []byte

//go:embed source/dedibox.yaml
var dediboxData []byte

type DataReader struct {
	ElasticMetalData map[string]model.Server
	DediboxData      map[string]model.Server
	AppleSiliconData map[string]model.Server
	InstancesData    map[string]model.VirtualMachine
}

func (d *DataReader) ReadElasticMetalData() error {
	err := yaml.Unmarshal(elasticMetalData, &d.ElasticMetalData)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataReader) ReadDediboxData() error {
	ddxMode := viper.GetString("dedibox.mode")
	if ddxMode == "on" {
		err := yaml.Unmarshal(dediboxData, &d.DediboxData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *DataReader) ReadAppleSiliconData() error {
	err := yaml.Unmarshal(appleSiliconData, &d.AppleSiliconData)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataReader) ReadInstancesData() error {
	err := yaml.Unmarshal(instancesData, &d.InstancesData)
	if err != nil {
		return err
	}

	return nil
}

func (d *DataReader) ReadVMsFromFile(filename string) (map[string]model.VirtualMachine, error) {
	fh, err := os.ReadFile(filename)
	if err != nil {
		log.Error("file read", log.WithError(err))
		return nil, err
	}

	result := make(map[string]model.VirtualMachine)

	err = yaml.Unmarshal(fh, &result)
	if err != nil {
		log.Error("file unmarshal", log.WithError(err))
		return nil, err
	}

	return result, nil
}

func (d *DataReader) ReadServersFromFile(filename string) (map[string]model.Server, error) {
	fh, err := os.ReadFile(filename)
	if err != nil {
		log.Error("file read", log.WithError(err))
		return nil, err
	}

	result := make(map[string]model.Server)

	err = yaml.Unmarshal(fh, &result)
	if err != nil {
		log.Error("file unmarshal", log.WithError(err))
		return nil, err
	}

	return result, nil
}

func NewDataReader() (*DataReader, error) {
	reader := DataReader{}

	err := reader.ReadAppleSiliconData()
	if err != nil {
		return nil, err
	}

	err = reader.ReadDediboxData()
	if err != nil {
		return nil, err
	}

	err = reader.ReadInstancesData()
	if err != nil {
		return nil, err
	}

	err = reader.ReadElasticMetalData()
	if err != nil {
		return nil, err
	}

	return &reader, nil
}
