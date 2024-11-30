package model

import "fmt"

const (
	ProductDedibox      = "DDX"
	ProductElasticMetal = "EM"
	ProductAppleSilicon = "AS"
)

type Server struct {
	Product string
	Name    string

	Cpus []Cpu
	Gpus []Gpu
	Rams []Ram
	Ssds []Ssd
	Hdds []Hdd

	VCpuPerCpuUnit uint32

	Motherboard Motherboard
	PowerSupply PowerSupply
}

func (s *Server) CpuName() string {
	if len(s.Cpus) == 0 {
		return ""
	}

	return s.Cpus[0].Name
}

func (s *Server) CpuCores() uint32 {
	if len(s.Cpus) == 0 {
		return 0
	}

	return s.Cpus[0].CoreUnits
}

func (s *Server) TotalCpuUnits() uint32 {
	units := uint32(0)
	for _, cpu := range s.Cpus {
		units += cpu.Units
	}

	return units
}

func (s *Server) TotalCores() uint32 {
	cores := uint32(0)
	for _, cpu := range s.Cpus {
		cores += cpu.CoreUnits
	}

	return cores
}

func (s *Server) GpuName() string {
	if len(s.Gpus) == 0 {
		return ""
	}

	return s.Gpus[0].Name
}

func (s *Server) TotalGpuUnits() uint32 {
	units := uint32(0)
	for _, gpu := range s.Gpus {
		units += gpu.Units
	}

	return units
}

func (s *Server) TotalGpuMemory() uint32 {
	memory := uint32(0)
	for _, gpu := range s.Gpus {
		memory += gpu.MemoryMib
	}

	return memory
}

func (s *Server) TotalRamUnits() uint32 {
	units := uint32(0)
	for _, ram := range s.Rams {
		units += ram.Units
	}

	return units
}

func (s *Server) TotalRamCapacity() uint32 {
	capacity := uint32(0)
	for _, ram := range s.Rams {
		capacity += ram.CapacityMib
	}

	return capacity
}

func (s *Server) TotalSsdUnits() uint32 {
	units := uint32(0)
	for _, ssd := range s.Ssds {
		units += ssd.Units
	}

	return units
}

func (s *Server) TotalSsdCapacity() uint32 {
	capacity := uint32(0)
	for _, ssd := range s.Ssds {
		capacity += ssd.CapacityMib
	}

	return capacity
}

func (s *Server) TotalHddUnits() uint32 {
	units := uint32(0)
	for _, hdd := range s.Hdds {
		units += hdd.Units
	}

	return units
}

func (s *Server) TotalHddCapacity() uint32 {
	capacity := uint32(0)
	for _, hdd := range s.Hdds {
		capacity += hdd.CapacityMib
	}

	return capacity
}

func DefaultServer(vCpus uint32, ramGiB uint32) Server {
	return Server{
		Cpus:           []Cpu{DefaultCpu(1, vCpus)},
		Rams:           []Ram{DefaultRam(1, ramGiB*1024)},
		VCpuPerCpuUnit: vCpus * 2,
		PowerSupply:    DefaultPowerSupply(500),
	}
}

func ServerToString(server Server) string {
	cores := uint32(0)
	ram := uint32(0)
	ssd := uint32(0)
	hdd := uint32(0)
	cpuName := "CPU"

	if len(server.Cpus) > 0 {
		cpuName = server.Cpus[0].Name
	}

	for _, c := range server.Cpus {
		cores += c.CoreUnits * c.Units
	}

	for _, r := range server.Rams {
		ram += r.CapacityMib * r.Units
	}

	for _, s := range server.Ssds {
		ssd += s.CapacityMib * s.Units
	}

	for _, h := range server.Hdds {
		hdd += h.CapacityMib * h.Units
	}

	res := fmt.Sprintf("%v core %v, %v GiB RAM", cores, cpuName, ram/1024)

	if ssd > 0 {
		res = fmt.Sprintf("%v, %v GiB SSD", res, ssd/1024)
	}

	if hdd > 0 {
		res = fmt.Sprintf("%v, %v GiB HDD", res, hdd/1024)
	}

	return res
}
