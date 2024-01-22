package model

func DefaultPowerSupply(units int32) PowerSupply {
	return PowerSupply{
		Units:    2,
		WeightKg: 2,
	}
}

type PowerSupply struct {
	Model        string
	Manufacturer string
	Units        int32
	WeightKg     int32
}
