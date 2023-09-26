# Data

To improve accuracy of server measurements, we need to provde more data and fill in some unknowns.

The defaults are configured in the respective files in the `model` directory, e.g. `model/instances.go`.

_VMs_

- Instance per server - this defines how many VMs are sharing a given host. This is just estimated at the moment. VM usages is calculated by dividing the impact of the underlying server by the number of VMs on the host

_CPUs_

- Type estimates - we are using the published "XYZ or equivalent" models for CPUs, and not specific models
- Thermal Design Power (TDP) - indicates how much heat a CPU/GPU generates under load. This is a number published by the manufacturer for each model, and we are just taking this at face value

_RAM_

- Model - we do not know the manufacturer of the RAM used in the underlying hosts so we are just using the default
- Capacity - this seems to be the only major factor on energy consumption, and we are guessing the RAM layout

_Disk_

- Manufacturer - as with RAM, the manufacturer of the underlying disks is not made available
