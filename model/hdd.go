package model

type Hdd struct {
	CapacityMB uint32
	Units      uint32
}

func DefaultHdds(units uint32, capacityMB uint32) []Hdd {
	return []Hdd{
		{
			CapacityMB: capacityMB,
			Units:      units,
		},
	}
}
