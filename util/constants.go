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
	InstanceZones     = []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3}
)
