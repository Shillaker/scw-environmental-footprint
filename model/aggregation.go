package model

func CombineImpacts(impacts []ImpactServerUsage) ImpactServerUsage {
	var aggregated ImpactServerUsage
	aggregated.Impacts = make(map[string]Impact)

	for _, i := range impacts {
		for name, impact := range i.Impacts {
			if aggImpact, ok := aggregated.Impacts[name]; ok {
				aggImpact.Manufacture += impact.Manufacture
				aggImpact.Use += impact.Use

				aggregated.Impacts[name] = aggImpact
			} else {
				aggregated.Impacts[name] = Impact{
					Manufacture: impact.Manufacture,
					Use:         impact.Use,
					Unit:        impact.Unit,
				}
			}
		}
	}

	return aggregated
}
