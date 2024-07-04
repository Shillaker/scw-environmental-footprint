package model

func DefaultPowerSupply(powerWatts uint32) PowerSupply {
	return PowerSupply{
		Units:    1,
		Watts:    powerWatts,
		WeightKg: 2,
	}
}

type PowerSupply struct {
	Model        string
	Manufacturer string
	Units        uint32
	Watts        uint32
	WeightKg     uint32
}

func (p PowerSupply) YearlyConsumptionKwh() uint32 {
	return (24 * 365 * p.Watts) / 1000
}
