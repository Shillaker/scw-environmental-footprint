package mapping

import (
	"testing"

	"github.com/shillaker/scw-environmental-footprint/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstanceMapping(t *testing.T) {
	mapper, err := NewScwMapper()
	require.NoError(t, err)

	instance := model.Instance{
		Type: "pro2-l",
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
		Server:    model.BasePro2Host,
		Usage:     expectedUsage,
		HostShare: 0.125,
	}

	expectedServerUsages := []model.ServerUsage{expectedServerUsage, expectedServerUsage}

	serverUsages, err := mapper.MapInstanceUsage(instance, usage)
	require.NoError(t, err)
	assert.EqualValues(t, expectedServerUsages, serverUsages)
}
