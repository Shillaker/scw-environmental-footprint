package model

func DefaultRam(units uint32, capacityMiB uint32) Ram {
	return Ram{
		CapacityMib:  capacityMiB,
		Units:        units,
	}
}

func DefaultRams(units uint32, capacityMiB uint32) []Ram {
	return []Ram{
		{
			CapacityMib:  capacityMiB,
			Units:        units,
		},
	}
}

type Ram struct {
	Units        uint32
	FrequencyHz  uint32
	Type         string

	CapacityMib uint32
}
