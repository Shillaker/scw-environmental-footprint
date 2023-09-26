package model

import "fmt"

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
)

type Instance struct {
	Type        string
	Description string
}

type InstanceBaseServer struct {
	VCpus  int32
	RamGiB int32
	Server Server
}

func (i *InstanceBaseServer) GetHostShare() float32 {
	totalCores := int32(0)
	for _, cpu := range i.Server.Cpus {
		totalCores += cpu.CoreUnits * cpu.Units
	}

	return float32(totalCores) / float32(i.VCpus)
}

// PLAY2 - 32 cores, AMD EPYC 7543, 64GiB
var BasePlay2Host = Server{
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Rams: DefaultRams(4, 16*1024),
}

// PRO2 - 32 cores, AMD EPYC 7543, 128GiB
var BasePro2Host = Server{
	Cpus: []Cpu{
		AmdEpyc7543,
	},
	Rams: DefaultRams(4, 32*1024),
}

// DEV1 - 16 cores, AMD EPYC 7281, 32GiB, 20GiB SSD
var BaseDev1Host = Server{
	Cpus: []Cpu{
		AmdEpyc7281,
	},
	Rams: DefaultRams(2, 16*1024),
	Ssds: []Ssd{
		{
			Manufacturer: ManufacturerMicron,
			CapacityMib:  20 * 1024,
			Units:        1,
		},
	},
}

// GP1 - 48 cores, AMD EPYC 7281, 256GiB, 600GiB SSD
var BaseGp1Host = Server{
	Cpus: []Cpu{
		AmdEpyc7281Cores48,
	},
	Rams: DefaultRams(8, 32*1024),
	Ssds: []Ssd{
		{
			Manufacturer: ManufacturerMicron,
			CapacityMib:  600 * 1024,
			Units:        1,
		},
	},
}

// POP2 - 64 cores, AMD EPYC 7543, 256GiB
var BasePop2Host = Server{
	Cpus: []Cpu{
		AmdEpyc7543Cores64,
	},
	Rams: DefaultRams(8, 32*1024),
}

// POP2HM - 64 cores, AMD EPYC 7543, 512GiB
var BasePop2HmHost = Server{
	Cpus: []Cpu{
		AmdEpyc7543Cores64,
	},
	Rams: DefaultRams(16, 32*1024),
}

func InstanceToString(instanceBase InstanceBaseServer) string {
	cpuModel := instanceBase.Server.Cpus[0].Model
	return fmt.Sprintf("%v vCPU, %v GiB, %v CPU", instanceBase.VCpus, instanceBase.RamGiB, cpuModel)
}

func buildInstanceBase(baseServer Server, nVcpu int32, ramGiB int32) InstanceBaseServer {
	server := baseServer

	return InstanceBaseServer{
		VCpus:  nVcpu,
		RamGiB: ramGiB,
		Server: server,
	}
}

var InstanceServerMapping = map[string]InstanceBaseServer{
	InstancePlay2Pico:  buildInstanceBase(BasePlay2Host, 1, 2),
	InstancePlay2Nano:  buildInstanceBase(BasePlay2Host, 2, 4),
	InstancePlay2Micro: buildInstanceBase(BasePlay2Host, 4, 8),
	InstancePro2Xxs:    buildInstanceBase(BasePro2Host, 2, 8),
	InstancePro2Xs:     buildInstanceBase(BasePro2Host, 4, 16),
	InstancePro2S:      buildInstanceBase(BasePro2Host, 8, 32),
	InstancePro2M:      buildInstanceBase(BasePro2Host, 16, 64),
	InstancePro2L:      buildInstanceBase(BasePro2Host, 32, 128),
	InstanceDev1S:      buildInstanceBase(BaseDev1Host, 2, 2),
	InstanceDev1M:      buildInstanceBase(BaseDev1Host, 3, 4),
	InstanceDev1L:      buildInstanceBase(BaseDev1Host, 4, 8),
	InstanceDev1Xl:     buildInstanceBase(BaseDev1Host, 4, 12),
	InstanceGp1Xs:      buildInstanceBase(BaseGp1Host, 4, 16),
	InstanceGp1S:       buildInstanceBase(BaseGp1Host, 8, 32),
	InstanceGp1M:       buildInstanceBase(BaseGp1Host, 16, 64),
	InstanceGp1L:       buildInstanceBase(BaseGp1Host, 32, 128),
	InstanceGp1Xl:      buildInstanceBase(BaseGp1Host, 48, 256),
	InstancePop22c:     buildInstanceBase(BasePop2Host, 2, 8),
	InstancePop24c:     buildInstanceBase(BasePop2Host, 4, 16),
	InstancePop28c:     buildInstanceBase(BasePop2Host, 8, 32),
	InstancePop16c:     buildInstanceBase(BasePop2Host, 16, 64),
	InstancePop32c:     buildInstanceBase(BasePop2Host, 32, 128),
	InstancePop64c:     buildInstanceBase(BasePop2Host, 64, 256),
	InstancePop2HM2C:   buildInstanceBase(BasePop2HmHost, 2, 16),
	InstancePop2HM4C:   buildInstanceBase(BasePop2HmHost, 4, 32),
	InstancePop2HM8C:   buildInstanceBase(BasePop2HmHost, 8, 64),
	InstancePop2HM16C:  buildInstanceBase(BasePop2HmHost, 16, 128),
	InstancePop2HM32C:  buildInstanceBase(BasePop2HmHost, 32, 256),
	InstancePop2HM64C:  buildInstanceBase(BasePop2HmHost, 64, 512),
}
