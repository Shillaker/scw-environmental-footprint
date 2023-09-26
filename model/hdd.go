package model

type Hdd struct {
	Model        string
	Manufacturer string

	CapacityMib int32
	Units       int32
}

func DefaultHdds(units int32, capacityMiB int32) []Hdd {
	return []Hdd{
		{
			CapacityMib: capacityMiB,
			Units:       units,
		},
	}
}
