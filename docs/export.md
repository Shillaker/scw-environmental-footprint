# Data export

To run the export, you will need to set up your local dev environment as described in [the development doc](./development.md)].

## Boavizta

The Boavizta export needs to output two files:

1. Instances - containing the virtualized resources available to each instance type, and the type of underlying server they run on
2. Servers - the specs for the underlying servers the instances run on

To generate these files, you can run:

```
make boavizta-export
```

The outputs can then be found in `outputs/`, as `instances.csv` and `servers.csv`.
