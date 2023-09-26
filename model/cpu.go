package model

const (
	ManufacturerIntel = "intel"
	ManufacturerAmd   = "amd"

	CpuMicroArchIntelIvybridge = "ivybridge"
	CpuMicroArchIntelSkylake   = "skylake"

	CpuFamilyAmdEpyc  = "epyc"
	CpuFamilyAmdRyzen = "ryzen"

	DefaultCpuManufacturer = ManufacturerIntel
	DefaultCpuMicroArch    = CpuMicroArchIntelSkylake
	DefaultCpuTdp          = 150
)

func DefaultCpu(units int32, cores int32) Cpu {
	return Cpu{
		Manufacturer: DefaultCpuManufacturer,
		CoreUnits:    cores,
		Tdp:          DefaultCpuTdp,
		Family:       DefaultCpuMicroArch,
		Units:        units,
	}
}

type Cpu struct {
	Model        string
	Manufacturer string
	Units        int32

	CoreUnits int32
	Tdp       int32
	Family    string
	Name      string
}

// https://www.amd.com/en/product/10941
var AmdEpyc7543 = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdEpyc,
	Model:        "EPYC 7543",
	CoreUnits:    32,
	Tdp:          225,
	Units:        1,
}

// https://www.amd.com/en/product/2001
var AmdEpyc7543Cores64 = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdEpyc,
	Model:        "EPYC 7543",
	CoreUnits:    32,
	Tdp:          225,
	Units:        2,
}

// https://www.amd.com/en/product/2001
var AmdEpyc7281 = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdEpyc,
	Model:        "EPYC 7281",
	CoreUnits:    16,
	Tdp:          155,
	Units:        1,
}

// https://www.amd.com/en/product/2001
var AmdEpyc7281Cores48 = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdEpyc,
	Model:        "EPYC 7281",
	CoreUnits:    16,
	Tdp:          155,
	Units:        3,
}

// https://www.amd.com/en/product/8866
var AmdRyzenPro3600 = Cpu{
	Manufacturer: ManufacturerAmd,
	Family:       CpuFamilyAmdRyzen,
	Model:        "RYZEN PRO 3600",
	CoreUnits:    6,
	Tdp:          65,
	Units:        1,
}

// https://ark.intel.com/content/www/us/en/ark/products/75777/intel-xeon-processor-e51410-v2-10m-cache-2-80-ghz.html
var IntelXeonE51410V2 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 1410 V2",
	CoreUnits:    4,
	Tdp:          80,
	Units:        1,
}

// https://ark.intel.com/content/www/us/en/ark/products/75780/intel-xeon-processor-e51650-v2-12m-cache-3-50-ghz.html
var IntelXeonE51650V2 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 1650 V2",
	CoreUnits:    6,
	Tdp:          130,
	Units:        1,
}

// https://www.intel.com/content/www/us/en/products/sku/75789/intel-xeon-processor-e52620-v2-15m-cache-2-10-ghz/specifications.html
var IntelXeonE52620V2 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 2620 V2",
	CoreUnits:    6,
	Tdp:          80,
	Units:        2,
}

//
var IntelXeonE52670 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 2670",
	CoreUnits:    8,
	Tdp:          115,
	Units:        2,
}

// https://ark.intel.com/content/www/us/en/ark/products/75275/intel-xeon-processor-e52670-v2-25m-cache-2-50-ghz.html
var IntelXeonE52670V2 = Cpu{
	Manufacturer: ManufacturerIntel,
	Family:       CpuMicroArchIntelIvybridge,
	Model:        "Xeon E5 2670 V2",
	CoreUnits:    10,
	Tdp:          115,
	Units:        2,
}
