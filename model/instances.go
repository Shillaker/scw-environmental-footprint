package model

import "fmt"

// An identifier for an instance type
type Instance struct {
	Type        string
	Description string
}

// Definition of a VM, including the server it runs on
type VirtualMachine struct {
	Type   string
	VCpus  uint32
	RamGiB uint32
	HddGiB uint32
	SsdGiB uint32
	Gpus   uint32
	Server Server
}

func DefaultInstanceSsd(capacityGiB uint32) []Ssd {
	return []Ssd{
		{
			CapacityMib: capacityGiB * 1024,
			Units:       1,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	}
}

// The percentage share of the impact of the underlying host attributable to the instance
func (i *VirtualMachine) GetHostShare() float32 {
	totalVCpus := uint32(0)

	for _, cpu := range i.Server.Cpus {
		if i.Server.VCpuPerCpu == 0 {
			totalVCpus += cpu.Threads // Assume dedicated if we don't know otherwise
		} else {
			totalVCpus += cpu.Units * i.Server.VCpuPerCpu
		}
	}

	return float32(i.VCpus) / float32(totalVCpus)
}

// Base for the Play2 range (shared vCPUs)
var BasePlay2Host = Server{
	Name: "play2.base",
	Cpus: []Cpu{
		AmdEpyc7543Double,
	},
	VCpuPerCpu: 2 * AmdEpyc7543.Threads,
	Rams: []Ram{
		{
			CapacityMib: 64 * 1024,
			Units:       16,
			Type:        "ddr4",
			FrequencyHz: 3200 * 1000 * 1000,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMib: 480 * 1024,
			Units:       2,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units:    2,
		Watts:    800,
		WeightKg: DefaultPowerSupplyWeightKg,
	},
}

// Base for the PRO2 range (shared vCPUs)
var BasePro2Host = Server{
	Name: "pro2.base",
	Cpus: []Cpu{
		AmdEpyc7543Double,
	},
	VCpuPerCpu: 2 * AmdEpyc7543.Threads,
	Rams: []Ram{
		{
			CapacityMib: 32 * 1024,
			Units:       24,
			Type:        "ddr4",
			FrequencyHz: 3200 * 1000 * 1000,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMib: 480 * 1024,
			Units:       2,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units:    2,
		Watts:    800,
		WeightKg: DefaultPowerSupplyWeightKg,
	},
}

// Base for the DEV1 range (shared vCPUs)
var BaseDev1Host = Server{
	Name: "dev1.base",
	Cpus: []Cpu{
		AmdEpyc7281,
	},
	Rams: []Ram{
		{
			CapacityMib: 32 * 1024,
			Units:       8,
			Type:        "ddr4",
			FrequencyHz: 2666 * 1000 * 1000,
		},
	},
	VCpuPerCpu: 2 * AmdEpyc7281.Threads,
	Ssds: []Ssd{
		{
			CapacityMib: 1024 * 1024,
			Units:       5,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units:    2,
		Watts:    800,
		WeightKg: DefaultPowerSupplyWeightKg,
	},
}

// Base for the GP1 range (shared vCPUs)
var BaseGp1Host = Server{
	Name: "gp1.base",
	Cpus: []Cpu{
		AmdEpyc7401P,
	},
	VCpuPerCpu: 2 * AmdEpyc7401P.Threads,
	Rams: []Ram{
		{
			Units:       12,
			CapacityMib: 32 * 1024,
			Type:        "ddr4",
			FrequencyHz: 2666 * 1000 * 1000,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMib: 1024 * 1024,
			Units:       5,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units:    2,
		Watts:    800,
		WeightKg: DefaultPowerSupplyWeightKg,
	},
}

// Base for the POP2 range (dedicated vCPUs)
var BasePop2Host = Server{
	Name: "pop2.base",
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Ssds: []Ssd{
		{
			CapacityMib: 1024 * 1024,
			Units:       5,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	Rams:        DefaultRams(8, 32*1024),
	VCpuPerCpu:  AmdEpyc7543.Threads,
	PowerSupply: DefaultPowerSupply(400),
}

// Base for the POP2HM range (dedicated vCPUs)
var BasePop2HmHost = Server{
	Name: "pop2hm.base",
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Ssds: []Ssd{
		{
			CapacityMib: 1024 * 1024,
			Units:       5,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	Rams:        DefaultRams(16, 32*1024),
	VCpuPerCpu:  AmdEpyc7543.Threads,
	PowerSupply: DefaultPowerSupply(400),
}

// Base for the POP2HC range (dedicated vCPUs)
var BasePop2HcHost = Server{
	Name: "pop2hc.base",
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Ssds: []Ssd{
		{
			CapacityMib: 1024 * 1024,
			Units:       5,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	Rams:       DefaultRams(16, 32*1024),
	VCpuPerCpu: AmdEpyc7543.Threads,
	PowerSupply: DefaultPowerSupply(400),
}

// Base for the STARDUST1 range (shared vCPUs)
var BaseStardust1Host = Server{
	Name: "stardust1.base",
	Cpus: []Cpu{
		AmdEpyc7281,
	},
	VCpuPerCpu: 2 * AmdEpyc7281.Threads,
	Rams: []Ram{
		{
			Units:       8,
			CapacityMib: 32 * 1024,
			Type:        "ddr4",
			FrequencyHz: 2666 * 1000 * 1000,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMib: 1024 * 1024,
			Units:       5,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units:    2,
		Watts:    800,
		WeightKg: DefaultPowerSupplyWeightKg,
	},
}

// Base for the ENT1 range (dedicated vCPUs, or not?)
var BaseEnt1Host = Server{
	Name: "ent1.base",
	Cpus: []Cpu{
		AmdEpyc7543Double,
	},
	VCpuPerCpu: 2 * AmdEpyc7543Triple.Threads,
	Rams: []Ram{
		{
			Units:       16,
			CapacityMib: 32 * 1024,
			Type:        "ddr4",
			FrequencyHz: 3200 * 1000 * 1000,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMib: 240 * 1024,
			Units:       2,
			Technology:  SsdTechnologyMlc,
			Casing:      SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units:    2,
		Watts:    800,
		WeightKg: DefaultPowerSupplyWeightKg,
	},
}

// Base for the COP ARM1 range (shared vCPUs)
var BaseCopArm1Host = Server{
	Name: "coparm1.base",
	Cpus: []Cpu{
		AmpereAltraMaxM12832,
	},
	Rams:       DefaultRams(8, 16*1024),
	VCpuPerCpu: 2 * AmpereAltraMaxM12832.Threads,
}

// Convert an instance into a readable string representation
func InstanceToString(instanceBase VirtualMachine) string {
	cpuName := instanceBase.Server.Cpus[0].Name
	return fmt.Sprintf("%v vCPU, %v GiB, %v CPU", instanceBase.VCpus, instanceBase.RamGiB, cpuName)
}

var BaseInstanceServers = []Server{
	BasePlay2Host,
	BasePro2Host,
	BaseDev1Host,
	BaseEnt1Host,
	BaseCopArm1Host,
	BaseGp1Host,
	BasePop2Host,
	BasePop2HmHost,
	BasePop2HcHost,
	BaseStardust1Host,
}
