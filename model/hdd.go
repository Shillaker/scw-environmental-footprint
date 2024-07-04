package model

type Hdd struct {
	Model        string
	Manufacturer string

	CapacityMib uint32
	Units       uint32
}

func DefaultHdds(units uint32, capacityMiB uint32) []Hdd {
	return []Hdd{
		{
			CapacityMib: capacityMiB,
			Units:       units,
		},
	}
}
