# Data

## Generating source data

The source data is pulled from the [Scaleway API](https://www.scaleway.com/en/developers/api), using the [Scaleway Go SDK](https://www.scaleway.com/en/docs/developer-tools/scaleway-sdk/go-sdk/).

To update the data, you can do the following:

- Create a [Scaleway Console](https://console.scaleway.com) account, and an associated API key
- Install [direnv]()
- Add a `.envrc` in the root of this directory containing
  - `CARBON_SCW_ACCESS_KEY` - the access key for your API key
  - `CARBON_SCW_SECRET_KEY` - the secret key for your API key
- Run `direnv allow`
- Run `task update-scw`

## TODOs

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

_Kubernetes_

- Control plane - we assume that the control plane runs on a standard Scaleway instance, but this may not be true
