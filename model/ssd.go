package model

const (
	SsdTechnologyTlc = "tlc"
	SsdTechnologyQlc = "qlc"
	SsdTechnologySlc = "slc"
	SsdTechnologyMlc = "mlc"

	SsdCasingM2     = "casing_m2"
	SsdCasing25Inch = "casing_2_5inch"
)

func DefaultSsd(units uint32, capacityMiB uint32) Ssd {
	return Ssd{
		CapacityMib:  capacityMiB,
		Units:        units,
		Technology:   SsdTechnologyMlc,
		Casing:       SsdCasingM2,
	}
}

func DefaultSsds(units uint32, capacityMiB uint32) []Ssd {
	return []Ssd{
		{
			CapacityMib:  capacityMiB,
			Units:        units,
			Technology:   SsdTechnologyMlc,
			Casing:       SsdCasingM2,
		},
	}
}

type Ssd struct {
	Units        uint32
	Technology   string
	Casing       string

	CapacityMib uint32
}
