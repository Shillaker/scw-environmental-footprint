package util

import (
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var initialised bool

func InitConfig() error {
	if initialised {
		return nil
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	configFileLocation := filepath.Join(userHomeDir, ".config", "scw", "carbon.yml")
	log.Infof("using config file: %v", configFileLocation)

	viper.SetConfigFile(configFileLocation)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Have viper automatically override config if it finds the appropriate environment variable
	// E.g. CARBON_GATEWAY.BACKEND_HOST will override the corresponding key in the config file
	viper.SetEnvPrefix("CARBON")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	// Set up defaults for docker compose environment
	viper.SetDefault("gateway.port", 8083)
	viper.SetDefault("gateway.backend_host", "backend")
	viper.SetDefault("gateway.backend_port", 8082)
	viper.SetDefault("impact.mode", "boavizta")

	viper.SetDefault("boavizta.host", "boavizta")
	viper.SetDefault("boavizta.port", 5000)

	viper.SetDefault("resilio.base_url", "https://db.resilio.tech/api")
	viper.SetDefault("resilio.token", "foobar")

	initialised = true

	return nil
}
