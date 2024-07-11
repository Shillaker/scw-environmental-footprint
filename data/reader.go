package data

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/shillaker/scw-environmental-footprint/model"
)

type DataReader struct {
}

func (d *DataReader) ReadElasticMetalData(ctx context.Context) (map[string]model.Server, error) {
	return d.ReadServersFromFile(ctx, EmFile)
}

func (d *DataReader) ReadDediboxData(ctx context.Context) (map[string]model.Server, error) {
	return d.ReadServersFromFile(ctx, DediboxFile)
}

func (d *DataReader) ReadAppleSiliconData(ctx context.Context) (map[string]model.Server, error) {
	return d.ReadServersFromFile(ctx, AppleSiliconFile)
}

func (d *DataReader) ReadInstancesData(ctx context.Context) (map[string]model.VirtualMachine, error) {
	return d.ReadVMsFromFile(ctx, InstancesFile)
}

func (d *DataReader) ReadVMsFromFile(ctx context.Context, filename string) (map[string]model.VirtualMachine, error) {
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

func (d *DataReader) ReadServersFromFile(ctx context.Context, filename string) (map[string]model.Server, error) {
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
