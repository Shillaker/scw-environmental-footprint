package model

type ElasticMetal struct {
	Type        string
	Description string
}

const (
	ElasticMetalA210 = "a210r-hdd"
	ElasticMetalA315 = "a315x-ssd"
	ElasticMetalA410 = "a410x-ssd"

	ElasticMetalB111X = "b111x-sata"
	ElasticMetalB112X = "b112x-ssd"
	ElasticMetalB211X = "b211x-sata"
	ElasticMetalB212X = "b212x-ssd"
	ElasticMetalB311X = "b311x-sata"
	ElasticMetalB312X = "b312x-ssd"

	ElasticMetalI210E = "ir-210e-nvme"

	ElasticMetalL101X = "li-101x-sata"
	ElasticMetalL105X = "li-105x-sata"
	ElasticMetalL110X = "li-110x-sata"

	ElasticMetalT210e = "t-t210e-nvme"
	ElasticMetalT510x = "t-t510x-nvme"
)

var ElasticMetalServerMapping = map[string]Server{
	ElasticMetalA210: {
		Cpus: []Cpu{AmdRyzenPro3600},
		Rams: DefaultRams(2, 8*1024),
		Hdds: DefaultHdds(2, 1024*1024),
	},
	ElasticMetalA315: {
		Cpus: []Cpu{IntelXeonE51410V2},
		Rams: DefaultRams(2, 32*1024),
		Ssds: DefaultSsds(2, 1024*1024),
	},
	ElasticMetalA410: {
		Cpus: []Cpu{IntelXeonE51650V2},
		Rams: DefaultRams(2, 32*1024),
		Ssds: DefaultSsds(2, 1024*1024),
	},
	ElasticMetalB111X: {
		Cpus: []Cpu{IntelXeonE52620V2},
		Rams: DefaultRams(6, 32*1024),
		Hdds: DefaultHdds(2, 8*1024*1024),
	},
	ElasticMetalB112X: {
		Cpus: []Cpu{IntelXeonE52620V2},
		Rams: DefaultRams(6, 32*1024),
		Ssds: DefaultSsds(2, 1024*1024),
	},
	ElasticMetalB211X: {
		Cpus: []Cpu{IntelXeonE52670},
		Rams: DefaultRams(8, 32*1024),
		Hdds: DefaultHdds(2, 8*1024*1024),
	},
	ElasticMetalB212X: {
		Cpus: []Cpu{IntelXeonE52670},
		Rams: DefaultRams(8, 32*1024),
		Ssds: DefaultSsds(2, 1024*1024),
	},
	ElasticMetalB311X: {
		Cpus: []Cpu{IntelXeonE52670V2},
		Rams: DefaultRams(8, 32*1024),
		Hdds: DefaultHdds(2, 12*1024*1024),
	},
	ElasticMetalB312X: {
		Cpus: []Cpu{IntelXeonE52670V2},
		Rams: DefaultRams(8, 32*1024),
		Ssds: DefaultSsds(2, 1024*1024),
	},
}
