package boavizta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.infra.online.net/paas/carbon/model"
)

func TestBoaviztaImpact(t *testing.T) {
	calculator, err := NewBoaviztaImpactCalculator()
	assert.NoError(t, err)

	serverA := model.Server{
		Cpus: []model.Cpu{
			model.DefaultCpu(2, 16*1024),
		},
		Rams: []model.Ram{
			model.DefaultRam(4, 32*1024),
		},
		Ssds: []model.Ssd{
			model.DefaultSsd(2, 256*1024),
		},
		Motherboard: model.DefaultMotherboard(1),
		PowerSupply: model.DefaultPowerSupply(1),
	}

	usageYears := 1
	usageSeconds := usageYears * 3600 * 24 * 365
	usageA := model.ServerUsageAmount{
		TimeSeconds:    int32(usageSeconds),
		LifespanYears:  4,
		LoadPercentage: 50,
		Region:         model.RegionFrance,
	}

	serverUsage := []model.ServerUsage{
		{Server: serverA, Usage: usageA, HostShare: 1},
	}

	// Set up expectations
	adpManufacture := float32(0.032)
	adpUse := float32(0.000131)
	adpUnit := "kgSbeq"

	gwpManufacture := float32(170)
	gwpUse := float32(260)
	gwpUnit := "kgCO2eq"

	peManufacture := float32(2300)
	peUse := float32(30430)
	peUnit := "MJ"

	// Calculate the impact
	impact, err := calculator.CalculateServerImpact(serverUsage)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(impact.Impacts))

	// Iterate over all keys
	for key, val := range impact.Impacts {
		if key == "adp" {
			assert.Equal(t, adpManufacture, val.Manufacture)
			assert.Equal(t, adpUse, val.Use)
			assert.Equal(t, adpUnit, val.Unit)
		} else if key == "gwp" {
			assert.Equal(t, gwpManufacture, val.Manufacture)
			assert.Equal(t, gwpUse, val.Use)
			assert.Equal(t, gwpUnit, val.Unit)
		} else if key == "pe" {
			assert.Equal(t, peManufacture, val.Manufacture)
			assert.Equal(t, peUse, val.Use)
			assert.Equal(t, peUnit, val.Unit)
		} else {
			assert.Fail(t, "Unrecognised impact type")
		}
	}

	// Check equivalents
	assert.Equal(t, impact.EquivalentsManufacture[0].Thing, model.EquivalentFlightLondonNY)
	assert.Equal(t, impact.EquivalentsUse[0].Thing, model.EquivalentFlightLondonNY)
}
