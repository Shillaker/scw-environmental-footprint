package main

import (
	"context"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/shillaker/scw-environmental-footprint/data"
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

	writer, err := data.NewDataWriter(ctx)
	if err != nil {
		log.Fatal("failed to init config", log.WithError(err))
	}

	writer.WriteAllProductData(ctx)
}
