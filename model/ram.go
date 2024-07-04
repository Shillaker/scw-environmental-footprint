package model

const (
	ManufacturerSamsung    = "samsung"
	DefaultRamManufacturer = ManufacturerSamsung
)

func DefaultRam(units uint32, capacityMiB uint32) Ram {
	return Ram{
		CapacityMib:  capacityMiB,
		Manufacturer: DefaultRamManufacturer,
		Units:        units,
	}
}

func DefaultRams(units uint32, capacityMiB uint32) []Ram {
	return []Ram{
		{
			CapacityMib:  capacityMiB,
			Manufacturer: DefaultRamManufacturer,
			Units:        units,
		},
	}
}

type Ram struct {
	Manufacturer string
	Units        uint32
	FrequencyHz  uint32
	Type         string

	CapacityMib uint32
}
