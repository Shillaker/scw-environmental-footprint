package data

import "path/filepath"

var (
	DataSourceDir    = filepath.Join("data", "source")
	DediboxFile      = filepath.Join(DataSourceDir, "dedibox.yaml")
	EmFile           = filepath.Join(DataSourceDir, "elastic_metal.yaml")
	AppleSiliconFile = filepath.Join(DataSourceDir, "apple_silicon.yaml")
	InstancesFile    = filepath.Join(DataSourceDir, "instances.yaml")
)
