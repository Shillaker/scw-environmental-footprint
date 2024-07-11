package mapping

import (
	"testing"

	"github.com/shillaker/scw-environmental-footprint/model"
	"github.com/stretchr/testify/assert"
)

func TestInstanceMapping(t *testing.T) {
	mapper, err := NewScwMapper()

	assert.NoError(t, err)

	instance := model.Instance{
		Type: model.InstancePlay2Pico,
	}

	timeDays := 60
	timeSeconds := int32(timeDays * 24 * 60 * 60)

	usage := model.CloudUsageAmount{
		Count:          2,
		LoadPercentage: 35,
		TimeSeconds:    timeSeconds,
		Region:         model.RegionFrance,
	}

	expectedUsage := model.DefaultUsage(timeSeconds)
	expectedUsage.LoadPercentage = 35

	expectedServerUsage := model.ServerUsage{
		Server: model.Server{
			Cpus: []model.Cpu{
				model.AmdEpyc7543,
			},
			Rams: []model.Ram{
				{
					CapacityMib: 16 * 1024,
					Units:       4,
				},
			},
		},
		Usage:     expectedUsage,
		HostShare: 32,
	}

	expectedServerUsages := []model.ServerUsage{expectedServerUsage, expectedServerUsage}

	serverUsages, err := mapper.MapInstanceUsage(instance, usage)

	assert.NoError(t, err)
	assert.EqualValues(t, expectedServerUsages, serverUsages)
}
