package model

func DefaultMotherboard(units int32) Motherboard {
	return Motherboard{}
}

type Motherboard struct {
	Model        string
	Manufacturer string
}

