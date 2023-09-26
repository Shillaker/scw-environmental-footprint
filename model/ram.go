package model

const (
	ManufacturerSamsung    = "samsung"
	DefaultRamManufacturer = ManufacturerSamsung
)

func DefaultRam(units int32, capacityMiB int32) Ram {
	return Ram{
		CapacityMib:  capacityMiB,
		Manufacturer: DefaultRamManufacturer,
		Units:        units,
	}
}

func DefaultRams(units int32, capacityMiB int32) []Ram {
	return []Ram{
		{
			CapacityMib:  capacityMiB,
			Manufacturer: DefaultRamManufacturer,
			Units:        units,
		},
	}
}

type Ram struct {
	Model        string
	Manufacturer string
	Units        int32

	CapacityMib int32
}
