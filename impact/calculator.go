package impact

import (
	"gitlab.infra.online.net/paas/carbon/impact/boavizta"
	"gitlab.infra.online.net/paas/carbon/model"
	"gitlab.infra.online.net/paas/carbon/util"

	"github.com/spf13/viper"
)

type ImpactCalculator interface {
	CalculateServerImpact(serverUsage []model.ServerUsage) (model.ImpactServerUsage, error)
}

var currentCalculator ImpactCalculator

func GetCalculator() (ImpactCalculator, error) {
	if currentCalculator == nil {

		var err error

		impactMode := viper.GetString("impact.mode")

		switch impactMode {
		case "boavizta":
			currentCalculator, err = boavizta.NewBoaviztaImpactCalculator()
		default:
			err = util.ErrInvalidImpactMode
		}

		if err != nil {
			return nil, err
		}
	}

	return currentCalculator, nil
}
