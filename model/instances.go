package model

import "fmt"

// ----- Docs -----
// Base server types: https://www.scaleway.com/en/docs/compute/instances/reference-content/instances-datasheet/#developement-and-general-purpose-instances
// Instances table: https://www.scaleway.com/en/pricing/?tags=compute

const (
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

// Instance - an identifier for an instance type
type Instance struct {
	Type        string
	Description string
}

// InstanceBaseServer - definition of an instance's virtualized resources, and the server it runs on
type InstanceBaseServer struct {
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
			Manufacturer: ManufacturerMicron,
			CapacityMib:  capacityGiB * 1024,
			Units:        1,
			Technology:   SsdTechnologyMlc,
			Casing:       SsdCasingM2,
		},
	}
}

// GetHostShare - the percentage share of the impact of the underlying host attributable to the instance
func (i *InstanceBaseServer) GetHostShare() float32 {
	totalVCpus := uint32(0)

	for _, cpu := range i.Server.Cpus {
		totalVCpus += cpu.Units * i.Server.VCpuPerCpu
	}

	return float32(i.VCpus) / float32(totalVCpus)
}

// BasePlay2Host - base for the Play2 range (shared vCPUs)
var BasePlay2Host = Server{
	Name: "play2.base",
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Rams:        DefaultRams(4, 16*1024),
	VCpuPerCpu:  2 * AmdEpyc7543.Threads,
	PowerSupply: DefaultPowerSupply(400),
}

// BasePro2Hose - base for the PRO2 range (shared vCPUs)
var BasePro2Host = Server{
	Name: "pro2.base",
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Rams:        DefaultRams(4, 32*1024),
	VCpuPerCpu:  2 * AmdEpyc7543.Threads,
	PowerSupply: DefaultPowerSupply(400),
}

// BaseDev1Host - base for the DEV1 range (shared vCPUs)
var BaseDev1Host = Server{
	Name: "dev1.base",
	Cpus: []Cpu{
		AmdEpyc7281,
	},
	Rams:        DefaultRams(2, 16*1024),
	VCpuPerCpu:  2 * AmdEpyc7281.Threads,
	Ssds:        DefaultInstanceSsd(20),
	PowerSupply: DefaultPowerSupply(400),
}

// BaseGp1Host - base for the GP1 range (shared vCPUs)
var BaseGp1Host = Server{
	Name: "gp1.base",
	Cpus: []Cpu{
		AmdEpyc7401P,
	},
	Rams:        DefaultRams(8, 32*1024),
	VCpuPerCpu:  2 * AmdEpyc7401P.Threads,
	Ssds:        DefaultInstanceSsd(600),
	PowerSupply: DefaultPowerSupply(400),
}

// BasePop2Host - base for the POP2 range (dedicated vCPUs)
var BasePop2Host = Server{
	Name: "pop2.base",
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Rams:        DefaultRams(8, 32*1024),
	VCpuPerCpu:  AmdEpyc7543.Threads,
	PowerSupply: DefaultPowerSupply(400),
}

// BasePop2HmHost - base for the POP2HM range (dedicated vCPUs)
var BasePop2HmHost = Server{
	Name: "pop2hm.base",
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Rams:        DefaultRams(16, 32*1024),
	VCpuPerCpu:  AmdEpyc7543.Threads,
	PowerSupply: DefaultPowerSupply(400),
}

// BasePop2HcHost - base for the POP2HC range (dedicated vCPUs)
var BasePop2HcHost = Server{
	Name: "pop2hc.base",
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Rams:       DefaultRams(16, 32*1024),
	VCpuPerCpu: AmdEpyc7543.Threads,
}

// BaseStardust1Host - base for the STARDUST1 range (shared vCPUs)
var BaseStardust1Host = Server{
	Name: "stardust1.base",
	Cpus: []Cpu{
		AmdEpyc7281,
	},
	Rams: DefaultRams(2, 8*1024),
	Ssds: []Ssd{
		{
			Manufacturer: ManufacturerMicron,
			CapacityMib:  10 * 1024,
			Units:        1,
		},
	},
	VCpuPerCpu: 2 * AmdEpyc7281.Threads,
}

// BaseCopArm1Host - base for the COP ARM1 range (shared vCPUs)
var BaseCopArm1Host = Server{
	Name: "coparm1.base",
	Cpus: []Cpu{
		AmpereAltraMaxM12832,
	},
	Rams:       DefaultRams(8, 16*1024),
	VCpuPerCpu: 2 * AmpereAltraMaxM12832.Threads,
}

// BaseEnt1Host - base for the ENT1 range (dedicated vCPUs)
var BaseEnt1Host = Server{
	Name: "ent1.base",
	Cpus: []Cpu{
		AmdEpyc7543Triple,
	},
	Rams:       DefaultRams(12, 32*1024),
	VCpuPerCpu: AmdEpyc7543Triple.Threads,
}

// InstanceToString - convert an instance into a readable string representation
func InstanceToString(instanceBase InstanceBaseServer) string {
	cpuName := instanceBase.Server.Cpus[0].Name
	return fmt.Sprintf("%v vCPU, %v GiB, %v CPU", instanceBase.VCpus, instanceBase.RamGiB, cpuName)
}

func buildInstanceBase(baseServer Server, nVcpu uint32, ramGiB uint32) InstanceBaseServer {
	server := baseServer

	return InstanceBaseServer{
		VCpus:  nVcpu,
		RamGiB: ramGiB,
		Server: server,
	}
}

var InstanceServerMapping = map[string]InstanceBaseServer{
	InstancePlay2Pico:     buildInstanceBase(BasePlay2Host, 1, 2),
	InstancePlay2Nano:     buildInstanceBase(BasePlay2Host, 2, 4),
	InstancePlay2Micro:    buildInstanceBase(BasePlay2Host, 4, 8),
	InstancePro2Xxs:       buildInstanceBase(BasePro2Host, 2, 8),
	InstancePro2Xs:        buildInstanceBase(BasePro2Host, 4, 16),
	InstancePro2S:         buildInstanceBase(BasePro2Host, 8, 32),
	InstancePro2M:         buildInstanceBase(BasePro2Host, 16, 64),
	InstancePro2L:         buildInstanceBase(BasePro2Host, 32, 128),
	InstanceDev1S:         buildInstanceBase(BaseDev1Host, 2, 2),
	InstanceDev1M:         buildInstanceBase(BaseDev1Host, 3, 4),
	InstanceDev1L:         buildInstanceBase(BaseDev1Host, 4, 8),
	InstanceDev1Xl:        buildInstanceBase(BaseDev1Host, 4, 12),
	InstanceGp1Xs:         buildInstanceBase(BaseGp1Host, 4, 16),
	InstanceGp1S:          buildInstanceBase(BaseGp1Host, 8, 32),
	InstanceGp1M:          buildInstanceBase(BaseGp1Host, 16, 64),
	InstanceGp1L:          buildInstanceBase(BaseGp1Host, 32, 128),
	InstanceGp1Xl:         buildInstanceBase(BaseGp1Host, 48, 256),
	InstancePop22c:        buildInstanceBase(BasePop2Host, 2, 8),
	InstancePop24c:        buildInstanceBase(BasePop2Host, 4, 16),
	InstancePop28c:        buildInstanceBase(BasePop2Host, 8, 32),
	InstancePop16c:        buildInstanceBase(BasePop2Host, 16, 64),
	InstancePop32c:        buildInstanceBase(BasePop2Host, 32, 128),
	InstancePop64c:        buildInstanceBase(BasePop2Host, 64, 256),
	InstancePop2HM2C:      buildInstanceBase(BasePop2HmHost, 2, 16),
	InstancePop2HM4C:      buildInstanceBase(BasePop2HmHost, 4, 32),
	InstancePop2HM8C:      buildInstanceBase(BasePop2HmHost, 8, 64),
	InstancePop2HM16C:     buildInstanceBase(BasePop2HmHost, 16, 128),
	InstancePop2HM32C:     buildInstanceBase(BasePop2HmHost, 32, 256),
	InstancePop2HM64C:     buildInstanceBase(BasePop2HmHost, 64, 512),
	InstancePop2HC2C:      buildInstanceBase(BasePop2HcHost, 2, 4),
	InstancePop2HC4C:      buildInstanceBase(BasePop2HcHost, 4, 8),
	InstancePop2HC8C:      buildInstanceBase(BasePop2HcHost, 8, 16),
	InstancePop2HC16C:     buildInstanceBase(BasePop2HcHost, 16, 32),
	InstancePop2HC32C:     buildInstanceBase(BasePop2HcHost, 32, 62),
	InstancePop2HC64C:     buildInstanceBase(BasePop2HcHost, 64, 128),
	InstanceStardust1:     buildInstanceBase(BaseStardust1Host, 1, 1),
	InstanceCopArm2C8G:    buildInstanceBase(BaseCopArm1Host, 2, 8),
	InstanceCopArm4C16G:   buildInstanceBase(BaseCopArm1Host, 4, 16),
	InstanceCopArm8C32G:   buildInstanceBase(BaseCopArm1Host, 8, 32),
	InstanceCopArm16C64G:  buildInstanceBase(BaseCopArm1Host, 16, 64),
	InstanceCopArm32C128G: buildInstanceBase(BaseCopArm1Host, 32, 128),
	InstanceEnt1Xxs:       buildInstanceBase(BaseEnt1Host, 2, 8),
	InstanceEnt1Xs:        buildInstanceBase(BaseEnt1Host, 4, 16),
	InstanceEnt1S:         buildInstanceBase(BaseEnt1Host, 8, 32),
	InstanceEnt1M:         buildInstanceBase(BaseEnt1Host, 16, 64),
	InstanceEnt1L:         buildInstanceBase(BaseEnt1Host, 32, 128),
	InstanceEnt1Xl:        buildInstanceBase(BaseEnt1Host, 64, 256),
	InstanceEnt1Xxl:       buildInstanceBase(BaseEnt1Host, 96, 384),
}

var InstanceBaseServers = []Server{
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
