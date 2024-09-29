package util

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	InitConfig()

	// Check values from config files
	assert.Equal(t, "5000", viper.GetString("boavizta.port"))
	assert.Equal(t, "localhost", viper.GetString("gateway.backend_host"))

	// Override a value and check
	err := os.Setenv("SCW_IMPACT_BOAVIZTA_PORT", "1234")
	assert.NoError(t, err)
	assert.Equal(t, "1234", viper.GetString("boavizta.port"))
	assert.Equal(t, "localhost", viper.GetString("gateway.backend_host"))

	// Override another value and check
	err = os.Setenv("SCW_IMPACT_GATEWAY_BACKEND_HOST", "foobar")
	assert.NoError(t, err)
	assert.Equal(t, "1234", viper.GetString("boavizta.port"))
	assert.Equal(t, "foobar", viper.GetString("gateway.backend_host"))

	os.Clearenv()
}
