package util

import (
	"strings"

	"github.com/spf13/viper"
)

var initialised bool

func InitConfig() error {
	if initialised {
		return nil
	}

	// Have viper automatically override config if it finds the appropriate environment variable
	// E.g. SCW_IMPACT_GATEWAY.BACKEND_HOST will override the corresponding key in the config file
	viper.SetEnvPrefix("SCW_IMPACT")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	// Set up defaults for docker compose environment
	viper.SetDefault("gateway.port", 8083)
	viper.SetDefault("gateway.backend_host", "backend")
	viper.SetDefault("gateway.backend_port", 8082)

	viper.SetDefault("boavizta.host", "boavizta")
	viper.SetDefault("boavizta.port", 5000)

	viper.SetDefault("resilio.base_url", "https://db.resilio.tech/api")
	viper.SetDefault("resilio.token", "foobar")

	viper.SetDefault("dedibox.mode", "off")

	viper.SetDefault("global.project_root", "/app")

	initialised = true

	return nil
}
