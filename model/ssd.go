package model

const (
	ManufacturerMicron     = "micron"
	DefaultSsdManufacturer = ManufacturerMicron

	SsdTechnologyTlc = "tlc"
	SsdTechnologyQlc = "qlc"
	SsdTechnologySlc = "slc"
	SsdTechnologyMlc = "mlc"

	SsdCasingM2     = "casing_m2"
	SsdCasing25Inch = "casing_2_5inch"
)

func DefaultSsd(units int32, capacityMiB int32) Ssd {
	return Ssd{
		CapacityMib:  capacityMiB,
		Manufacturer: DefaultSsdManufacturer,
		Units:        units,
		Technology:   SsdTechnologyMlc,
		Casing:       SsdCasingM2,
	}
}

func DefaultSsds(units int32, capacityMiB int32) []Ssd {
	return []Ssd{
		{
			CapacityMib:  capacityMiB,
			Manufacturer: DefaultSsdManufacturer,
			Units:        units,
			Technology:   SsdTechnologyMlc,
			Casing:       SsdCasingM2,
		},
	}
}

type Ssd struct {
	Model        string
	Manufacturer string
	Units        int32
	Technology   string
	Casing       string

	CapacityMib int32
}
