package model

const (
	DefaultLifespanYears  = 4
	DefaultLoadPercentage = 50
	DefaultRegion         = RegionFrance

	RegionPoland      = "pol"
	RegionFrance      = "fra"
	RegionNetherlands = "ned"
)

func DefaultUsage(usageSeconds int32) ServerUsageAmount {
	return ServerUsageAmount{
		TimeSeconds:    int32(usageSeconds),
		LifespanYears:  4,
		LoadPercentage: 50,
		Region:         RegionFrance,
	}
}

// -----------------------------
// Cloud
// -----------------------------

type CloudUsageAmount struct {
	TimeSeconds    int32
	Count          int32
	LoadPercentage int32
	MemoryMiB      int32
	MilliVCPU      int32

	Region string
}

// -----------------------------
// Servers
// -----------------------------

type ServerUsageAmount struct {
	TimeSeconds    int32
	LifespanYears  int32
	LoadPercentage float32
	Region         string
}

func (s ServerUsageAmount) UsageAsPercentageOfLifespan() float32 {
	return float32(s.TimeSeconds) / float32(s.LifespanSeconds())
}

func (s ServerUsageAmount) TimeHoursRoundedUp() int32 {
	secondsInHour := int32(60 * 60)
	return int32((s.TimeSeconds + (secondsInHour - 1)) / secondsInHour)
}

func (s ServerUsageAmount) LifespanSeconds() int32 {
	return s.LifespanYears * 365 * 24 * 60 * 60
}

type ServerUsage struct {
	Server    Server
	Usage     ServerUsageAmount
	HostShare float32
}

type Impact struct {
	Manufacture float32
	Use         float32
	Unit        string
}

type EquivalentCO2E struct {
	Amount float32
	Thing  string
}

type ImpactServerUsage struct {
	Impacts                map[string]Impact
	EquivalentsManufacture []EquivalentCO2E
	EquivalentsUse         []EquivalentCO2E
}
