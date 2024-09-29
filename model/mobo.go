package model

func DefaultMotherboard(units int32) Motherboard {
	return Motherboard{}
}

type Motherboard struct {
	Units uint32
}
