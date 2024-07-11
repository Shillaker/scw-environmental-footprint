package model

import "fmt"

// ----- Docs -----
// Base server types: https://www.scaleway.com/en/docs/compute/instances/reference-content/instances-datasheet/#developement-and-general-purpose-instances
// Instances table: https://www.scaleway.com/en/pricing/?tags=compute

const (
	DefaultInstanceHddGib = 20

	InstancePlay2Pico  = "play2-pico"
	InstancePlay2Nano  = "play2-nano"
	InstancePlay2Micro = "play2-micro"

	InstancePro2Xxs = "pro2-xxs"
	InstancePro2Xs  = "pro2-xs"
	InstancePro2S   = "pro2-s"
	InstancePro2M   = "pro2-m"
	InstancePro2L   = "pro2-l"

	InstanceDev1S  = "dev1-s"
	InstanceDev1M  = "dev1-m"
	InstanceDev1L  = "dev1-l"
	InstanceDev1Xl = "dev1-xl"

	InstanceGp1Xs = "gp1-xs"
	InstanceGp1S  = "gp1-s"
	InstanceGp1M  = "gp1-m"
	InstanceGp1L  = "gp1-l"
	InstanceGp1Xl = "gp1-xl"

	InstancePop22c = "pop2-2c-8g"
	InstancePop24c = "pop2-4c-16g"
	InstancePop28c = "pop2-8c-32g"
	InstancePop16c = "pop2-16c-64g"
	InstancePop32c = "pop2-32c-128g"
	InstancePop64c = "pop2-64c-256g"

	InstanceEnt1Xxs = "ent1-xxs"
	InstanceEnt1Xs  = "ent1-xs"
	InstanceEnt1S   = "ent1-s"
	InstanceEnt1M   = "ent1-xm"
	InstanceEnt1L   = "ent1-l"
	InstanceEnt1Xl  = "ent1-xl"
	InstanceEnt1Xxl = "ent1-2xl"

	InstancePop2HM2C  = "pop2-hm-2c-16g"
	InstancePop2HM4C  = "pop2-hm-4c-32g"
	InstancePop2HM8C  = "pop2-hm-8c-64g"
	InstancePop2HM16C = "pop2-hm-16c-128g"
	InstancePop2HM32C = "pop2-hm-32c-256g"
	InstancePop2HM64C = "pop2-hm-64c-512g"

	InstancePop2HC2C  = "pop2-hc-2c-4g"
	InstancePop2HC4C  = "pop2-hc-4c-8g"
	InstancePop2HC8C  = "pop2-hc-8c-16g"
	InstancePop2HC16C = "pop2-hc-16c-32g"
	InstancePop2HC32C = "pop2-hc-32c-64g"
	InstancePop2HC64C = "pop2-hc-64c-128g"

	InstanceStardust1 = "stardust1"

	InstanceCopArm2C8G    = "coparm1-2c-8g"
	InstanceCopArm4C16G   = "coparm1-4c-16g"
	InstanceCopArm8C32G   = "coparm1-8c-32g"
	InstanceCopArm16C64G  = "coparm1-16c-64g"
	InstanceCopArm32C128G = "coparm1-32c-128g"
)

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

func buildInstanceBase(baseServer Server, nVcpu uint32, ramGiB uint32) VirtualMachine {
	server := baseServer

	return VirtualMachine{
		VCpus:  nVcpu,
		RamGiB: ramGiB,
		Server: server,
	}
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
