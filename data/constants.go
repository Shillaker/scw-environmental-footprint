package data

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func GetDataSoureDir() string {
	return filepath.Join(viper.GetString("global.project_root"), "data", "source")
}

func GetDediboxFile() string {
	return filepath.Join(GetDataSoureDir(), "dedibox.yaml")
}

func GetElasticMetalFile() string {
	return filepath.Join(GetDataSoureDir(), "elastic_metal.yaml")
}

func GetAppleSiliconFile() string {
	return filepath.Join(GetDataSoureDir(), "apple_silicon.yaml")
}

func GetInstancesFile() string {
	return filepath.Join(GetDataSoureDir(), "instances.yaml")
}
