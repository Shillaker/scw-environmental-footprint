package model

const (
	DefaultGpuName = "Nvidia H100"
)

func DefaultGpu(units uint32) Gpu {
	return Gpu{
		Name:  DefaultGpuName,
		Units: units,
	}
}

type Gpu struct {
	Manufacturer string
	Name         string
	Units        uint32
	MemoryMib    uint32
	Cores        uint32
}
