package model

import "fmt"

const (
	ManufacturerNvidia = "nvidia"

	DefaultGpuManufacturer = ManufacturerNvidia
)

func DefaultGpu(units int32) Gpu {
	return Gpu{
		Manufacturer: DefaultGpuManufacturer,
		Units:        units,
	}
}

type Gpu struct {
	Model        string
	Manufacturer string
	Units        int32

	MemoryMib int32
}

func (g Gpu) Name() string {
	return fmt.Sprintf("%s %s", g.Manufacturer, g.Model)
}
