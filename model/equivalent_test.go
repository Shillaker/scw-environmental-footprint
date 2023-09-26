package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroCO2EEquivalent(t *testing.T) {
	zeroCO2EKg := float32(0.0)

	expected := EquivalentCO2E{
		Thing:  EquivalentFlightLondonNY,
		Amount: 0.0,
	}
	actuals := CalculateEquivalentCO2E(zeroCO2EKg)
	assert.Equal(t, 6, len(actuals))

	actual := actuals[0]
	assert.Equal(t, expected, actual)
}

func TestCO2EEquivalent(t *testing.T) {
	co2EKgIn := float32(130.5)

	expected := EquivalentCO2E{
		Thing:  EquivalentFlightLondonNY,
		Amount: 0.145,
	}
	actuals := CalculateEquivalentCO2E(co2EKgIn)
	actual := actuals[0]
	assert.Equal(t, expected, actual)
}
