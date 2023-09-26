package model

const (
	ManufacturerMicron     = "micron"
	DefaultSsdManufacturer = ManufacturerMicron
)

func DefaultSsd(units int32, capacityMiB int32) Ssd {
	return Ssd{
		CapacityMib:  capacityMiB,
		Manufacturer: DefaultSsdManufacturer,
		Units:        units,
	}
}

func DefaultSsds(units int32, capacityMiB int32) []Ssd {
	return []Ssd{
		{
			CapacityMib:  capacityMiB,
			Manufacturer: DefaultSsdManufacturer,
			Units:        units,
		},
	}
}

type Ssd struct {
	Model        string
	Manufacturer string
	Units        int32

	CapacityMib int32
}
