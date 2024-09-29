package model

const (
	DefaultCpuName = "AMD EPYC 7543"
)

func DefaultCpu(units uint32, cores uint32) Cpu {
	return Cpu{
		Name:        DefaultCpuName,
		CoreUnits:   cores,
		Threads:     cores * 2,
		Units:       units,
		FrequencyHz: 2500,
	}
}

type Cpu struct {
	Name  string
	Units uint32

	CoreUnits   uint32
	Threads     uint32
	FrequencyHz uint32
}

// https://www.amd.com/en/product/10941
var AmdEpyc7543 = Cpu{
	Name:        "AMD EPYC 7543",
	CoreUnits:   32,
	Threads:     64,
	Units:       1,
	FrequencyHz: 2800,
}

var AmdEpyc7543Double = Cpu{
	Name:        "AMD EPYC 7543",
	CoreUnits:   32,
	Threads:     64,
	Units:       2,
	FrequencyHz: 2800,
}

// https://www.amd.com/en/product/10941
var AmdEpyc7543Triple = Cpu{
	Name:        "AMD EPYC 7543",
	CoreUnits:   32,
	Threads:     64,
	Units:       3,
	FrequencyHz: 2800,
}

// https://www.amd.com/en/product/2001
var AmdEpyc7281 = Cpu{
	Name:        "AMD EPYC 7281",
	CoreUnits:   16,
	Threads:     32,
	Units:       1,
	FrequencyHz: 2100,
}

// https://www.amd.com/fr/support/cpu/amd-epyc/amd-epyc-7001-series/amd-epyc-7401p
var AmdEpyc7401P = Cpu{
	Name:        "AMD EPYC 7401P",
	CoreUnits:   24,
	Threads:     48,
	Units:       2,
	FrequencyHz: 2000,
}

// https://www.amd.com/en/product/8866
var AmdRyzenPro3600 = Cpu{
	Name:        "AMD RYZEN PRO 3600",
	CoreUnits:   6,
	Threads:     12,
	Units:       1,
	FrequencyHz: 3600,
}

// https://ark.intel.com/content/www/us/en/ark/products/75777/intel-xeon-processor-e51410-v2-10m-cache-2-80-ghz.html
var IntelXeonE51410V2 = Cpu{
	Name:        "Intel Xeon E5 1410 V2",
	CoreUnits:   4,
	Threads:     8,
	Units:       1,
	FrequencyHz: 2800,
}

// https://ark.intel.com/content/www/us/en/ark/products/75780/intel-xeon-processor-e51650-v2-12m-cache-3-50-ghz.html
var IntelXeonE51650V2 = Cpu{
	Name:        "Intel Xeon E5 1650 V2",
	CoreUnits:   6,
	Threads:     12,
	Units:       1,
	FrequencyHz: 3500,
}

// https://ark.intel.com/content/www/us/en/ark/products/64594/intel-xeon-processor-e5-2620-15m-cache-2-00-ghz-7-20-gt-s-intel-qpi.html
var IntelXeonE52620V2 = Cpu{
	Name:        "Intel Xeon E5 2620 V2",
	CoreUnits:   6,
	Threads:     12,
	Units:       2,
	FrequencyHz: 2000,
}

// https://ark.intel.com/content/www/fr/fr/ark/products/64595/intel-xeon-processor-e5-2670-20m-cache-2-60-ghz-8-00-gt-s-intel-qpi.html
var IntelXeonE52670 = Cpu{
	Name:        "Intel Xeon E5 2670",
	CoreUnits:   8,
	Units:       2,
	FrequencyHz: 2600,
}

// https://ark.intel.com/content/www/us/en/ark/products/75275/intel-xeon-processor-e52670-v2-25m-cache-2-50-ghz.html
var IntelXeonE52670V2 = Cpu{
	Name:        "Intel Xeon E5 2670 V2",
	CoreUnits:   10,
	Units:       2,
	FrequencyHz: 2500,
}

// https://d1o0i0v5q5lp8h.cloudfront.net/ampere/live/assets/documents/Altra_Max_Rev_A1_PB_v1.00_20220331.pdf
var AmpereAltraMaxM12832 = Cpu{
	Name:        "Ampere Altra Max 128",
	CoreUnits:   128,
	Threads:     128,
	Units:       1,
	FrequencyHz: 3000,
}
