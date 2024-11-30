package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/shillaker/scw-environmental-footprint/data"
	"github.com/shillaker/scw-environmental-footprint/util"
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

	err = writer.WriteAllProductData(ctx)
	if err != nil {
		log.Fatal("failed to write product data", log.WithError(err))
	}
}
