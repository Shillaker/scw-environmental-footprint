package util

import (
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	DefaultPageSize    = 50
	DefaultZone        = scw.ZoneFrPar1
	DefaultDediboxZone = "fr-par-2"
	DefaultRegion      = scw.RegionFrPar
)

var (
	DediboxZones      = []scw.Zone{scw.ZoneFrPar1, scw.ZoneNlAms1, scw.ZonePlWaw1}
	ElasticMetalZones = []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1}
	AppleSiliconZones = []scw.Zone{scw.ZoneFrPar1}
)
