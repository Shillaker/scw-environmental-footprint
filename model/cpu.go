package model

import "fmt"

const (
	ManufacturerIntel  = "Intel"
	ManufacturerAmd    = "AMD"
	ManufacturerAmpere = "Ampere"

	CpuMicroArchIntelIvybridge = "ivybridge"
	CpuMicroArchIntelSkylake   = "skylake"

	CpuFamilyAmdEpyc  = "Epyc"
	CpuFamilyAmdRyzen = "Ryzen"

	CpuFamilyAmpereAltra = "Altra"

	DefaultCpuManufacturer = ManufacturerIntel
	DefaultCpuMicroArch    = CpuMicroArchIntelSkylake
	DefaultCpuTdp          = 150
)

func DefaultCpu(units int32, cores int32) Cpu {
	return Cpu{
		Manufacturer: DefaultCpuManufacturer,
		CoreUnits:    cores,
		Threads:      cores * 2,
		Tdp:          DefaultCpuTdp,
		Family:       DefaultCpuMicroArch,
		Units:        units,
		FrequencyHz:  2500,
	}
}

type Cpu struct {
	Model        string
	Manufacturer string
	Units        int32

	CoreUnits   int32
	Threads     int32
	Tdp         int32
	FrequencyHz int32
	Family      string
}

func (c Cpu) Name() string {
	return fmt.Sprintf("%s %s", c.Manufacturer, c.Model)
}

// https://www.amd.com/en/product/10941
var AmdEpyc7543 = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdEpyc,
	Model:        "EPYC 7543",
	CoreUnits:    32,
	Threads:      64,
	Tdp:          225,
	Units:        1,
	FrequencyHz:  2800,
}

// https://www.amd.com/en/product/10941
var AmdEpyc7543Triple = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdEpyc,
	Model:        "EPYC 7543",
	CoreUnits:    32,
	Threads:      64,
	Tdp:          225,
	Units:        3,
	FrequencyHz:  2800,
}

// https://www.amd.com/en/product/2001
var AmdEpyc7281 = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdEpyc,
	Model:        "EPYC 7281",
	CoreUnits:    16,
	Threads:      32,
	Tdp:          155,
	Units:        1,
	FrequencyHz:  2100,
}

// https://www.amd.com/fr/support/cpu/amd-epyc/amd-epyc-7001-series/amd-epyc-7401p
var AmdEpyc7401P = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdEpyc,
	Model:        "EPYC 7401P",
	CoreUnits:    24,
	Threads:      48,
	Tdp:          155,
	Units:        2,
	FrequencyHz:  2000,
}

// https://www.amd.com/en/product/8866
var AmdRyzenPro3600 = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdRyzen,
	Model:        "RYZEN PRO 3600",
	CoreUnits:    6,
	Threads:      12,
	Tdp:          65,
	Units:        1,
	FrequencyHz:  3600,
}

// https://ark.intel.com/content/www/us/en/ark/products/75777/intel-xeon-processor-e51410-v2-10m-cache-2-80-ghz.html
var IntelXeonE51410V2 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 1410 V2",
	CoreUnits:    4,
	Threads:      8,
	Tdp:          80,
	Units:        1,
	FrequencyHz:  2800,
}

// https://ark.intel.com/content/www/us/en/ark/products/75780/intel-xeon-processor-e51650-v2-12m-cache-3-50-ghz.html
var IntelXeonE51650V2 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 1650 V2",
	CoreUnits:    6,
	Threads:      12,
	Tdp:          130,
	Units:        1,
	FrequencyHz:  3500,
}

// https://ark.intel.com/content/www/us/en/ark/products/64594/intel-xeon-processor-e5-2620-15m-cache-2-00-ghz-7-20-gt-s-intel-qpi.html
var IntelXeonE52620V2 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 2620 V2",
	CoreUnits:    6,
	Threads:      12,
	Tdp:          95,
	Units:        2,
	FrequencyHz:  2000,
}

// https://ark.intel.com/content/www/fr/fr/ark/products/64595/intel-xeon-processor-e5-2670-20m-cache-2-60-ghz-8-00-gt-s-intel-qpi.html
var IntelXeonE52670 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 2670",
	CoreUnits:    8,
	Tdp:          115,
	Units:        2,
	FrequencyHz:  2600,
}

// https://ark.intel.com/content/www/us/en/ark/products/75275/intel-xeon-processor-e52670-v2-25m-cache-2-50-ghz.html
var IntelXeonE52670V2 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 2670 V2",
	CoreUnits:    10,
	Tdp:          115,
	Units:        2,
	FrequencyHz:  2500,
}

// https://d1o0i0v5q5lp8h.cloudfront.net/ampere/live/assets/documents/Altra_Max_Rev_A1_PB_v1.00_20220331.pdf
var AmpereAltraMaxM12832 = Cpu{
	Manufacturer: ManufacturerAmpere,
	Family:       CpuFamilyAmpereAltra,
	Model:        "Altra Max 128",
	CoreUnits:    128,
	Threads:      128,
	Tdp:          250,
	Units:        1,
	FrequencyHz:  3000,
}
