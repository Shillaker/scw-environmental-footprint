package model

const (
	DefaultPowerSupplyWeightKg = 2
)

func DefaultPowerSupply(powerWatts uint32) PowerSupply {
	return PowerSupply{
		Units:    1,
		Watts:    powerWatts,
		WeightKg: DefaultPowerSupplyWeightKg,
	}
}

type PowerSupply struct {
	Units        uint32
	Watts        uint32
	WeightKg     uint32
}

func (p PowerSupply) YearlyConsumptionKwh() uint32 {
	return (24 * 365 * p.Watts) / 1000
}
