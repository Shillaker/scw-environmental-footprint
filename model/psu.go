package model

func DefaultPowerSupply(powerWatts int32) PowerSupply {
	return PowerSupply{
		Units: 1,
		Watts: powerWatts,
	}
}

type PowerSupply struct {
	Model        string
	Manufacturer string
	Units        int32

	Watts int32
}

func (p PowerSupply) YearlyConsumptionKwh() int32 {
	return (24 * 365 * p.Watts) / 1000
}
