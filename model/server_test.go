package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerModel(t *testing.T) {
	// Empty server
	server := Server{}
	assert.Equal(t, "0 core CPU, 0 GiB RAM", ServerToString(server))

	// CPU and RAM
	server.Cpus = []Cpu{
		{CoreUnits: 10, Units: 1, Model: "foobar"},
		{CoreUnits: 5, Units: 2},
	}

	server.Rams = []Ram{
		{CapacityMib: 32 * 1024, Units: 2},
		{CapacityMib: 16 * 1024, Units: 1},
	}
	assert.Equal(t, "20 core foobar, 80 GiB RAM", ServerToString(server))

	// SSD
	server.Ssds = []Ssd{
		{CapacityMib: 512 * 1024, Units: 2},
		{CapacityMib: 256 * 1024, Units: 1},
	}
	assert.Equal(t, "20 core foobar, 80 GiB RAM, 1280 GiB SSD", ServerToString(server))
}
