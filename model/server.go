package model

import "fmt"

type Server struct {
	Cpus []Cpu
	Rams []Ram
	Ssds []Ssd
	Hdds []Hdd

	Motherboard Motherboard
	PowerSupply PowerSupply
}

func DefaultServer(vCpus int32, ramGiB int32) Server {
	return Server{
		Cpus: []Cpu{DefaultCpu(1, vCpus)},
		Rams: []Ram{DefaultRam(1, ramGiB*1024)},
	}
}

func ServerToString(server Server) string {
	cores := int32(0)
	ram := int32(0)
	ssd := int32(0)
	hdd := int32(0)
	cpuModel := "CPU"

	if len(server.Cpus) > 0 {
		cpuModel = server.Cpus[0].Model
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

	res := fmt.Sprintf("%v core %v, %v GiB RAM", cores, cpuModel, ram/1024)

	if ssd > 0 {
		res = fmt.Sprintf("%v, %v GiB SSD", res, ssd/1024)
	}

	if hdd > 0 {
		res = fmt.Sprintf("%v, %v GiB HDD", res, hdd/1024)
	}

	return res
}
