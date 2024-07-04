package main

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"github.com/shillaker/scw-environmental-footprint/util"
)

var (
	dataSourceDir    = filepath.Join("data", "source")
	dediboxFile      = filepath.Join(dataSourceDir, "dedibox.yaml")
	emFile           = filepath.Join(dataSourceDir, "elastic_metal.yaml")
	appleSiliconFile = filepath.Join(dataSourceDir, "apple_silicon.yaml")
	instancesFile    = filepath.Join(dataSourceDir, "instances.yaml")
)

func main() {
	err := util.InitConfig()
	if err != nil {
		log.Fatal("failed to init config", log.WithError(err))
	}

	util.InitLogging()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err = os.MkdirAll(dataSourceDir, os.ModePerm)
	if err != nil {
		log.Fatal("could not create output dir", log.WithError(err))
	}

	client, err := util.NewClient(ctx)

	if err != nil {
		log.Fatal("could not create scw client", log.WithError(err))
	}

	ddxMode := viper.GetString("dedibox.mode")
	if ddxMode == "on" {
		dediboxServers, err := client.ListDediboxServers(ctx)
		if err != nil {
			log.Fatal("could not list ddx servers", log.WithError(err))
		}

		dediboxData, err := yaml.Marshal(&dediboxServers)
		if err != nil {
			log.Fatal("dedibox marshal", log.WithError(err))
		}

		err = ioutil.WriteFile(dediboxFile, dediboxData, 0766)
		if err != nil {
			log.Fatal("dedibox write", log.WithError(err))
		}
	}

	emServers, err := client.ListElasticMetalServers(ctx)
	if err != nil {
		log.Fatal("could not list em servers", log.WithError(err))
	}

	emData, err := yaml.Marshal(emServers)
	if err != nil {
		log.Fatal("em marshal", log.WithError(err))
	}

	err = ioutil.WriteFile(emFile, emData, 0766)
	if err != nil {
		log.Fatal("em write", log.WithError(err))
	}

	asServers, err := client.ListAppleSiliconServers(ctx)
	if err != nil {
		log.Fatal("could not list as servers", log.WithError(err))
	}

	asData, err := yaml.Marshal(asServers)
	if err != nil {
		log.Fatal("as marshal", log.WithError(err))
	}

	err = ioutil.WriteFile(appleSiliconFile, asData, 0766)
	if err != nil {
		log.Fatal("as write", log.WithError(err))
	}
}
