package util

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogging() {
	log.SetFormatter(&log.JSONFormatter{})

	logLevel := viper.GetString("logging.level")
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
		log.Debug("logging set to debug")
	} else if logLevel == "trace" {
		log.SetLevel(log.TraceLevel)
		log.Trace("logging set to trace")
	}
}
