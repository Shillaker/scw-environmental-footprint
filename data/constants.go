package data

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func GetDataSourceDir() string {
	return filepath.Join(viper.GetString("global.project_root"), "data", "source")
}

func GetDediboxFile() string {
	return filepath.Join(GetDataSourceDir(), "dedibox.yaml")
}

func GetElasticMetalFile() string {
	return filepath.Join(GetDataSourceDir(), "elastic_metal.yaml")
}

func GetAppleSiliconFile() string {
	return filepath.Join(GetDataSourceDir(), "apple_silicon.yaml")
}

func GetInstancesFile() string {
	return filepath.Join(GetDataSourceDir(), "instances.yaml")
}
