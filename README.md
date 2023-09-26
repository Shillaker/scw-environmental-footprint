# Environmental impact PoC

This is an experiment to measure the environmental impact measurements for Scaleway products.

There is a more detailed [architecture doc](docs/architecture.md) which describes how the pieces fit together.

The project uses the [BoaviztAPI](https://github.com/Boavizta/boaviztapi) to measure the impact of each service. Boavizta provide their own front-end for the API via [Datavizta](https://dataviz.boavizta.org/serversimpact).

## Running instance

A running instange can be found [here](http://e367ffdd-6c2f-4f53-a7ec-7fb7893c4896.pub.instances.scw.cloud:8081/).

## Details

The architecture is split into three layers:

1. API (+ gateway) - allowing users to query the impact of their usage of cloud services
2. Mapping - maps the usage of cloud services to usage of underlying devices (i.e. servers)
3. Impact - calculates the impact of server usage (i.e. calls BoaviztAPI)
