package resilio

const (
	ResilioRegionFrance      = "France"
	ResilioRegionNetherlands = "Netherlands"
	ResilioRegionPoland      = "Poland"
)

type ResilioCpu struct {
	Name string `json:"name"`
	// DieSurfaceMmSq uint32  `json:"die_surface_mm2"`
	// LithoNm        uint32  `json:"litho_nm"`
}

type ResilioRam struct {
	SizeGb uint32 `json:"size_gb"`
}

type ResilioSsd struct {
	SizeGb     uint32  `json:"size_gb"`
	Technology string `json:"technology"`
	Casing     string `json:"casing"`
}

type ResilioHddConfig struct {
	Quantity uint32 `json:"quantity"`
}

type ResilioGpu struct {
	Name           string `json:"name"`
	DieSurfaceMmSq uint32  `json:"die_surface_mm2"`
	LithoNm        uint32  `json:"litho_nm"`
}

type ResilioPsu struct {
	PowerWatt uint32 `json:"power_watt"`
}

type ResilioPowerUsage struct {
	YearlyElectricityConsumption uint32  `json:"yearly_electricity_consumption"`
	PowerWatt                    uint32  `json:"power_watt"`
	DeltaTHour                   uint32  `json:"delta_t_hour"`
	Geography                    string `json:"geography"`
}

type ResilioHoursUsage struct {
	DeltaTHour uint32 `json:"delta_t_hour"`
}

type ResilioRackServerUsage struct {
	UsagePercent      uint32             `json:"usage_percent"`
	WantedName        string            `json:"wanted_name"`
	RackUnit          uint32             `json:"rack_unit"`
	ContigurationName string            `json:"configuration_name"`
	Cpus              []ResilioCpu      `json:"cpus"`
	Rams              []ResilioRam      `json:"rams"`
	Ssds              []ResilioSsd      `json:"ssd_disks"`
	Hdds              ResilioHddConfig  `json:"hdd_disks"`
	Gpus              []ResilioGpu      `json:"dedicated_graphics_cards"`
	Psus              []ResilioPsu      `json:"power_supplies"`
	Usage             ResilioPowerUsage `json:"usage"`
}

type ResilioVmUsage struct {
	SizeRamGb  uint32                  `json:"size_ram_gb"`
	Pue        float32                `json:"pue"`
	Mirroring  float32                `json:"mirroring"`
	RackServer ResilioRackServerUsage `json:"rack_server"`
	Usage      ResilioHoursUsage      `json:"usage"`
}

type ResilioRackServerRequest struct {
	Assembly   bool                     `json:"assembly"`
	ServerData []ResilioRackServerUsage `json:"data"`
}

type ResilioVmRequest struct {
	Assembly   bool             `json:"assembly"`
	ServerData []ResilioVmUsage `json:"data"`
}

type ResilioImpact struct {
	ADPe   float32 `json:"ADPe"`
	ADPf   float32 `json:"ADPf"`
	AP     float32 `json:"AP"`
	CTUe   float32 `json:"CTUe"`
	CTUhc  float32 `json:"CTUhc"`
	CTUhnc float32 `json:"CTUhnc"`
	Epf    float32 `json:"Epf"`
	Epm    float32 `json:"Epm"`
	Ept    float32 `json:"Ept"`
	GWP    float32 `json:"GWP"`
	IR     float32 `json:"IR"`
	LU     float32 `json:"LU"`
	MIPS   float32 `json:"MIPS"`
	ODP    float32 `json:"ODP"`
	PM     float32 `json:"PM"`
	POCP   float32 `json:"POCP"`
	TPE    float32 `json:"TPE"`
	WU     float32 `json:"WU"`
}

type ResilioResultPerLcStep struct {
	Build        ResilioImpact `json:"BLD"`
	Distribution ResilioImpact `json:"DIS"`
	Usage        ResilioImpact `json:"USE"`
	EndOfLife    ResilioImpact `json:"EOL"`
}

func (r ResilioResultPerLcStep) GWPEmbedded() float32 {
	return r.Build.GWP + r.Distribution.GWP + r.EndOfLife.GWP
}

func (r ResilioResultPerLcStep) GWPUse() float32 {
	return r.Usage.GWP
}

func (r ResilioResultPerLcStep) ADPEmbedded() float32 {
	return r.Build.ADPe + r.Distribution.ADPe + r.EndOfLife.ADPe + r.Build.ADPf + r.Distribution.ADPf + r.EndOfLife.ADPf
}

func (r ResilioResultPerLcStep) ADPUse() float32 {
	return r.Usage.ADPe + r.Usage.ADPf
}

func (r ResilioResultPerLcStep) PEEmbedded() float32 {
	return r.Build.TPE + r.Distribution.TPE + r.EndOfLife.TPE
}

func (r ResilioResultPerLcStep) PEUse() float32 {
	return r.Usage.TPE
}

type ResilioImpacts struct {
	Total               ResilioImpact          `json:"total"`
	PerLcStep           ResilioResultPerLcStep `json:"per_lc_step"`
	NormalizedPerLcStep ResilioResultPerLcStep `json:"normalized_per_lc_step"`
}

type ResilioResults struct {
	Inventory ResilioImpacts `json:"inventory{}"`
	TryApi    ResilioImpacts `json:"Try API"`
}

type ResilioServerResponse struct {
	Results ResilioResults `json:"results"`
}
