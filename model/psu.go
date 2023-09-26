package model

func DefaultPowerSupply(units int32) PowerSupply {
	return PowerSupply{
		Units: 1,
	}
}

type PowerSupply struct {
	Model        string
	Manufacturer string
	Units        int32
}

