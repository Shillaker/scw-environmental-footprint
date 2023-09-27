# Environmental impact PoC

This is an experiment to measure the environmental impact measurements for Scaleway products.

The project uses the [BoaviztAPI](https://github.com/Boavizta/boaviztapi) to measure the impact of each service. Boavizta provide their own front-end for the API via [Datavizta](https://dataviz.boavizta.org/serversimpact).

## Running instance

A running instance can be found [here](http://scw-impact.simonshillaker.com/).

## Details

The architecture is split into three layers:

1. API (+ gateway) - allowing users to query the impact of their usage of cloud services
2. Mapping - maps the usage of cloud services to usage of underlying devices (i.e. servers)
3. Impact - calculates the impact of server usage (i.e. calls BoaviztAPI)
