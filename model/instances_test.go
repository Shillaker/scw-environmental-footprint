package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceHostShare(t *testing.T) {
	server := Server{
		Cpus: []Cpu{
			{CoreUnits: 32, Units: 2},
		},
	}

	instanceBase := VirtualMachine{
		VCpus:  4,
		Server: server,
	}

	assert.Equal(t, float32(16), instanceBase.GetHostShare())
}
