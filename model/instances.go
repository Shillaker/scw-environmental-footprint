package model

import "fmt"

// ----- Docs -----
// Instance types: https://www.scaleway.com/en/docs/compute/instances/reference-content/instances-datasheet/#developement-and-general-purpose-instances
// GPU instances: https://www.scaleway.com/en/docs/compute/gpu/reference-content/choosing-gpu-instance-type/
const (
	MinimumInstanceBlockVolume = 10
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
	HddGb  uint32
	SsdGb  uint32
	Gpus   uint32
	Server Server
}

func DefaultInstanceSsd(capacityGiB uint32) []Ssd {
	return []Ssd{
		{
			CapacityMB: capacityGiB * 1000,
			Units:      1,
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
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

func giveServerNewName(other Server, newName string) Server {
	new := other
	new.Name = newName
	return new
}

// Base for the Play2 range (shared vCPUs)
var BasePlay2Host = Server{
	Name: "scw_play2.base",
	Cpus: []Cpu{
		AmdEpyc7543Double,
	},
	VCpuPerCpu: 2 * AmdEpyc7543.Threads,
	Rams: []Ram{
		{
			CapacityMib: 64 * 1024, // 64GiB
			Units:       16,
			Type:        "ddr4",
			FrequencyHz: 3200e6,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMB: 480e3, // 480GB
			Units:      2,
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 2,
		Watts: 1400,
	},
}

// Base for the PRO2 range (shared vCPUs)
var BasePro2Host = Server{
	Name: "scw_pro2.base",
	Cpus: []Cpu{
		AmdEpyc7543Double,
	},
	VCpuPerCpu: 2 * AmdEpyc7543.Threads,
	Rams: []Ram{
		{
			CapacityMib: 32 * 1024, // 32GiB
			Units:       24,
			Type:        "ddr4",
			FrequencyHz: 3200e6,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMB: 480e3, // 480GiB
			Units:      2,
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 2,
		Watts: 1400,
	},
}

// Base for the DEV1 range (shared vCPUs)
var BaseDev1Host = Server{
	Name: "scw_dev1.base",
	Cpus: []Cpu{
		AmdEpyc7281,
	},
	Rams: []Ram{
		{
			CapacityMib: 32 * 1024, // 32GiB
			Units:       8,
			Type:        "ddr4",
			FrequencyHz: 2666e6,
		},
	},
	VCpuPerCpu: 2 * AmdEpyc7281.Threads,
	Ssds: []Ssd{
		{
			CapacityMB: 1e6, // 1TB
			Units:      5,
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 2,
		Watts: 500,
	},
}

// Base for the GP1 range (shared vCPUs)
var BaseGp1Host = Server{
	Name: "scw_gp1.base",
	Cpus: []Cpu{
		AmdEpyc7401P,
	},
	VCpuPerCpu: 2 * AmdEpyc7401P.Threads,
	Rams: []Ram{
		{
			Units:       12,
			CapacityMib: 32 * 1024, // 32GiB
			Type:        "ddr4",
			FrequencyHz: 2666e6,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMB: 1e6, // 1TB
			Units:      5,
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 2,
		Watts: 500,
	},
}

// Base for the ENT1 range (dedicated vCPUs, or not?)
var BaseEnt1Host = Server{
	Name: "scw_ent1.base",
	Cpus: []Cpu{
		AmdEpyc7543Double,
	},
	VCpuPerCpu: 2 * AmdEpyc7543Triple.Threads,
	Rams: []Ram{
		{
			Units:       16,
			CapacityMib: 32 * 1024, // 32GiB
			Type:        "ddr4",
			FrequencyHz: 3200e6,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMB: 240e3, // 240GB
			Units:      2,
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 2,
		Watts: 1400,
	},
}

// Base for the POP2 range (dedicated vCPUs)
var BasePop2Host = giveServerNewName(BaseEnt1Host, "scw_pop2.base")

// Base for the POP2HM range (dedicated vCPUs)
var BasePop2HmHost = giveServerNewName(BaseEnt1Host, "scw_pop2hm.base")

// Base for the POP2HC range (dedicated vCPUs)
var BasePop2HcHost = giveServerNewName(BaseEnt1Host, "scw_pop2hc.base")

// Base for the STARDUST1 range (shared vCPUs)
var BaseStardust1Host = Server{
	Name: "scw_stardust1.base",
	Cpus: []Cpu{
		AmdEpyc7281,
	},
	VCpuPerCpu: 2 * AmdEpyc7281.Threads,
	Rams: []Ram{
		{
			Units:       8,
			CapacityMib: 32 * 1024, // 32GiB
			Type:        "ddr4",
			FrequencyHz: 2666e6,
		},
	},
	Ssds: []Ssd{
		{
			CapacityMB: 1e6, // 1TB
			Units:      5,
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 2,
		Watts: 800,
	},
}

// Base for the H100 range
var BaseH100Host = Server{
	Name: "scw_h100.base",
	Cpus: []Cpu{
		AmdEpyc9334,
	},
	VCpuPerCpu: AmdEpyc9334.Threads,
	Rams: []Ram{
		{
			Units:       12,
			CapacityMib: 64 * 1024, // 64GiB
			Type:        "ddr5",
			FrequencyHz: 2400e6,
		},
	},
	Gpus: []Gpu{
		{
			Units:     2,
			Name:      "Nvidia H100",
			MemoryMib: 80 * 1024, // 80 GiB
		},
	},
	Ssds: []Ssd{
		{
			Units:      5,
			CapacityMB: 7.68e6, // 7.68TB
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 2,
		Watts: 1200,
	},
}

// Base for the L4 range
var BaseL4Host = Server{
	Name: "scw_l4.base",
	Cpus: []Cpu{
		AmdEpyc7413,
	},
	VCpuPerCpu: AmdEpyc7413.Threads,
	Rams: []Ram{
		{
			Units:       16,
			CapacityMib: 32 * 1024, // 32GiB
			Type:        "ddr4",
			FrequencyHz: 3200e6,
		},
	},
	Gpus: []Gpu{
		{
			Units:     8,
			Name:      "Nvidia L4",
			MemoryMib: 24 * 1024, // 24GiB
		},
	},
	Ssds: []Ssd{
		{
			Units:      2,
			CapacityMB: 240e3, // 240GB
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 4,
		Watts: 2000,
	},
}

// Base for the RenderS range
var BaseRenderSHost = Server{
	Name: "scw_renders.base",
	Cpus: []Cpu{
		IntelXeonGold6148,
	},
	VCpuPerCpu: IntelXeonGold6148.Threads,
	Rams: []Ram{
		{
			Units:       12,
			CapacityMib: 32 * 1024, // 32GiB
			Type:        "ddr4",
			FrequencyHz: 3200e6,
		},
	},
	Gpus: []Gpu{
		{
			Units:     8,
			Name:      "Nvidia Tesla P100",
			MemoryMib: 24 * 1024, // 24GiB
		},
	},
	Ssds: []Ssd{
		{
			Units:      2,
			CapacityMB: 3.84e6, // 3.84TB
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 4,
		Watts: 1600,
	},
}

// Base for the COP ARM1 range (shared vCPUs)
var BaseCopArm1Host = Server{
	Name: "scw_coparm1.base",
	Cpus: []Cpu{
		AmpereAltraMaxM12832,
	},
	Rams: []Ram{
		{
			Units:       16,
			CapacityMib: 64 * 1024, // 64GiB
			Type:        "ddr4",
			FrequencyHz: 2666e6,
		},
	},
	VCpuPerCpu: 2 * AmpereAltraMaxM12832.Threads,
	Ssds: []Ssd{
		{
			CapacityMB: 960e3, // 960GB
			Units:      2,
			Technology: SsdTechnologyMlc,
			Casing:     SsdCasingM2,
		},
	},
	PowerSupply: PowerSupply{
		Units: 2,
		Watts: 800,
	},
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
	BaseH100Host,
	BaseL4Host,
	BaseRenderSHost,
}
