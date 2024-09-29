# Boavizta

## Generating the Boavizta export

To run the export, you will need to set up your local dev environment as described in [the development doc](./development.md)].

The Boavizta export creates two files:

1. Instances - containing the virtualized resources available to each instance type, and the type of underlying server they run on
2. Servers - the specs for the underlying servers the instances run on

To generate these files, you can run:

```
task export-boavizta
```

The outputs can then be found in `data/output`, as `instances.csv` and `servers.csv`.

## Boavizta API in this project

This project uses the Boavizta API to esimate the impacts of Scaleway products.

This is done via two inputs:

1. Server/device configuration - CPU, RAM, SSD etc.
2. Usage - region, time, load

These are sent to the Boavizta `server` API.

## Links

### Docs

Useful links from the [BoaviztAPI docs](https://doc.api.boavizta.org/):

- [Usage methodology](https://doc.api.boavizta.org/Explanations/usage/usage/)
- [Manufacture methodology](https://doc.api.boavizta.org/Explanations/manufacture_methodology/)
- [Electrical impact factors](https://doc.api.boavizta.org/Explanations/usage/elec_factors/)
- [Impact criteria](https://doc.api.boavizta.org/Explanations/impacts/)
- [Components](https://doc.api.boavizta.org/Explanations/components/component/)
