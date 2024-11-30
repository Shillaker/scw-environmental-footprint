package boavizta

// Boavizta country codes: https://doc.api.boavizta.org/Explanations/usage/countries/
const (
	BoaviztaRegionFrance      = "FRA"
	BoaviztaRegionPoland      = "POL"
	BoaviztaRegionNetherlands = "NLD"

	// Allocation type is either LINEAR or TOTAL. LINEAR divides the manufacture impact over the usage time proportional to the lifespan
	BoaviztaAllocationTypeLinear = "LINEAR"
)

type BoaviztaCpu struct {
	Units     int32  `json:"units"`
	CoreUnits int32  `json:"core_units"`
	Family    string `json:"family"`
	Tdp       int32  `json:"tdp"`
}

type BoaviztaPowerSupply struct {
	Units int32 `json:"units"`
}

type BoaviztaRam struct {
	Units        int32  `json:"units"`
	CapacityGib  int32  `json:"capacity"`
	Manufacturer string `json:"manufacturer"`
}

type BoaviztaDisk struct {
	Type         string `json:"type"`
	CapacityGib  int32  `json:"capacity"`
	Manufacturer string `json:"manufacturer"`
	Units        int32  `json:"units"`
}

type BoaviztaServerConfiguration struct {
	Cpu         BoaviztaCpu         `json:"cpu"`
	Ram         []BoaviztaRam       `json:"ram"`
	Disk        []BoaviztaDisk      `json:"disk"`
	PowerSupply BoaviztaPowerSupply `json:"power_supply"`
}

type BoaviztaServerModel struct {
	Type string `json:"type"`
}

// The time workload is used to specify the level of utilisation of the device, in order to calculate energy consumption from the consumption profile
// https://doc.api.boavizta.org/Explanations/usage/elec_conso
type BoaviztaTimeWorkload struct {
	LoadPercentage float32 `json:"load_percentage"`
	TimePercentage float32 `json:"time_percentage"`
}

// See https://doc.api.boavizta.org/Reference/format/usage/
type BoaviztaServerUsage struct {
	UsageLocation string                 `json:"usage_location"`
	TimeWorkload  []BoaviztaTimeWorkload `json:"time_workload"`
}

type BoaviztaServerRequest struct {
	Configuration BoaviztaServerConfiguration `json:"configuration"`
	Model         BoaviztaServerModel         `json:"model"`
	Usage         BoaviztaServerUsage         `json:"usage"`
}

type BoaviztaImpact struct {
	Embedded    BoaviztaImpactValue `json:"embedded"`
	Use         BoaviztaImpactValue `json:"use"`
	Description string              `json:"description"`
	Unit        string              `json:"unit"`
}

type BoaviztaImpactValue struct {
	Value float32 `json:"value"`
	Min   float32 `json:"min"`
	Max   float32 `json:"max"`
}

type BoaviztaImpacts struct {
	GwpImpact BoaviztaImpact `json:"gwp"`
	PeImpact  BoaviztaImpact `json:"pe"`
	AdpImpact BoaviztaImpact `json:"adp"`
}

type BoaviztaServerResponse struct {
	Impacts BoaviztaImpacts `json:"impacts"`
}
