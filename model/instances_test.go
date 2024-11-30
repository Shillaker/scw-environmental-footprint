package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceHostShare(t *testing.T) {
	t.Run("with vCPU per CPU set", func(t *testing.T) {
		server := Server{
			Cpus: []Cpu{
				{CoreUnits: 32, Threads: 64, Units: 2},
			},
			VCpuPerCpuUnit: 128,
		}

		instanceBase := VirtualMachine{
			VCpus:  16,
			Server: server,
		}

		assert.Equal(t, float32(0.0625), instanceBase.GetHostShare())
	})

	t.Run("without vCPU per CPU set", func(t *testing.T) {
		server := Server{
			Cpus: []Cpu{
				{CoreUnits: 32, Threads: 64, Units: 2},
			},
		}

		instanceBase := VirtualMachine{
			VCpus:  16,
			Server: server,
		}

		assert.Equal(t, float32(0.25), instanceBase.GetHostShare())
	})

	t.Run("with multiple CPUs", func(t *testing.T) {
		server := Server{
			Cpus: []Cpu{
				{CoreUnits: 32, Threads: 64, Units: 2},
				{CoreUnits: 32, Threads: 64, Units: 2},
			},
		}

		instanceBase := VirtualMachine{
			VCpus:  16,
			Server: server,
		}

		assert.Equal(t, float32(0.125), instanceBase.GetHostShare())
	})
}
