package util

import (
	log "github.com/sirupsen/logrus"
)

func InitLogging() {
	log.SetFormatter(&log.JSONFormatter{})
}
