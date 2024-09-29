package data

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"github.com/shillaker/scw-environmental-footprint/model"
)

type DataReader struct {
	ElasticMetalData map[string]model.Server
	DediboxData      map[string]model.Server
	AppleSiliconData map[string]model.Server
	InstancesData    map[string]model.VirtualMachine
}

func (d *DataReader) ReadElasticMetalData() error {
	data, err := d.ReadServersFromFile(GetElasticMetalFile())
	if err != nil {
		return err
	}
	d.ElasticMetalData = data

	return nil
}

func (d *DataReader) ReadDediboxData() error {
	ddxMode := viper.GetString("dedibox.mode")
	if ddxMode == "on" {

		data, err := d.ReadServersFromFile(GetDediboxFile())
		if err != nil {
			return err
		}
		d.DediboxData = data
	}

	return nil
}

func (d *DataReader) ReadAppleSiliconData() error {
	data, err := d.ReadServersFromFile(GetAppleSiliconFile())
	if err != nil {
		return err
	}
	d.AppleSiliconData = data

	return nil
}

func (d *DataReader) ReadInstancesData() error {
	data, err := d.ReadVMsFromFile(GetInstancesFile())
	if err != nil {
		return err
	}
	d.InstancesData = data

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
