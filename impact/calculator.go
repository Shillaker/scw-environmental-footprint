package impact

import (
	"github.com/shillaker/scw-environmental-footprint/impact/boavizta"
	"github.com/shillaker/scw-environmental-footprint/impact/resilio"
	"github.com/shillaker/scw-environmental-footprint/model"
	"github.com/shillaker/scw-environmental-footprint/util"
)

type ImpactCalculator interface {
	CalculateServerImpact(serverUsage []model.ServerUsage) (model.ImpactServerUsage, error)
}

var (
	boaviztaCalculator *boavizta.BoaviztaImpactCalculator
	resilioCalculator  *resilio.ResilioImpactCalculator
)

func GetCalculator(config model.ImpactConfig) (ImpactCalculator, error) {
	var err error

	isBoavizta := config.Backend == "boavizta" || len(config.Backend) == 0
	isResilio := config.Backend == "resilio"

	if isBoavizta {
		if boaviztaCalculator == nil {
			boaviztaCalculator, err = boavizta.NewBoaviztaImpactCalculator()
			if err != nil {
				return nil, err
			}
		}

		return boaviztaCalculator, nil
	} else if isResilio {
		if resilioCalculator == nil {
			resilioCalculator, err = resilio.NewResilioImpactCalculator()
			if err != nil {
				return nil, err
			}
		}

		return resilioCalculator, nil
	} else {
		return nil, util.ErrInvalidImpactMode
	}
}
