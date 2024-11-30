package data

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"github.com/shillaker/scw-environmental-footprint/util"
)

type DataWriter struct {
	cli *util.SCWClient
}

func (d *DataWriter) WriteAllProductData(ctx context.Context) error {
	err := d.WriteDediboxData(ctx)
	if err != nil {
		return err
	}

	err = d.WriteAppleSiliconData(ctx)
	if err != nil {
		return err
	}

	err = d.WriteElasticMetalData(ctx)
	if err != nil {
		return err
	}

	err = d.WriteInstancesData(ctx)
	if err != nil {
		return err
	}

	return err
}

func (d *DataWriter) WriteDediboxData(ctx context.Context) error {
	ddxMode := viper.GetString("dedibox.mode")
	if ddxMode == "on" {
		dediboxServers, err := d.cli.ListDediboxServers(ctx)
		if err != nil {
			log.Error("could not list ddx servers", log.WithError(err))
			return err
		}

		dediboxData, err := yaml.Marshal(&dediboxServers)
		if err != nil {
			log.Error("dedibox marshal", log.WithError(err))
			return err
		}

		err = os.WriteFile(GetDediboxFile(), dediboxData, 0766)
		if err != nil {
			log.Error("dedibox write", log.WithError(err))
			return err
		}
	}

	return nil
}

func (d *DataWriter) WriteElasticMetalData(ctx context.Context) error {
	emServers, err := d.cli.ListElasticMetalServers(ctx)
	if err != nil {
		log.Error("could not list em servers", log.WithError(err))
		return err
	}

	emData, err := yaml.Marshal(emServers)
	if err != nil {
		log.Error("em marshal", log.WithError(err))
		return err
	}

	err = os.WriteFile(GetElasticMetalFile(), emData, 0766)
	if err != nil {
		log.Error("em write", log.WithError(err))
		return err
	}

	return nil
}

func (d *DataWriter) WriteAppleSiliconData(ctx context.Context) error {
	asServers, err := d.cli.ListAppleSiliconServers(ctx)
	if err != nil {
		log.Error("could not list as servers", log.WithError(err))
		return err
	}

	asData, err := yaml.Marshal(asServers)
	if err != nil {
		log.Error("as marshal", log.WithError(err))
		return err
	}

	err = os.WriteFile(GetAppleSiliconFile(), asData, 0766)
	if err != nil {
		log.Error("as write", log.WithError(err))
		return err
	}

	return nil
}

func (d *DataWriter) WriteInstancesData(ctx context.Context) error {
	instanceServers, err := d.cli.ListInstanceVMs(ctx)
	if err != nil {
		log.Error("could not list instance VMs", log.WithError(err))
		return err
	}

	instancesData, err := yaml.Marshal(instanceServers)
	if err != nil {
		log.Error("instance marshal", log.WithError(err))
		return err
	}

	err = os.WriteFile(GetInstancesFile(), instancesData, 0766)
	if err != nil {
		log.Error("instance write", log.WithError(err))
		return err
	}

	return nil
}

func NewDataWriter(ctx context.Context) (*DataWriter, error) {
	err := os.MkdirAll(GetDataSourceDir(), os.ModePerm)
	if err != nil {
		log.Error("could not create output dir", log.WithError(err))
		return nil, err
	}

	client, err := util.NewClient(ctx)

	if err != nil {
		log.Error("could not create scw client", log.WithError(err))
		return nil, err
	}

	return &DataWriter{
		cli: client,
	}, nil
}
