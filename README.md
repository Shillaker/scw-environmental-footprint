# 🌱 Scaleway environmental impact PoC

This is an experiment to measure the environmental impact measurements for Scaleway products.

You can see a running instance [here](https://scw-impact.simonshillaker.com/).

The project uses the [BoaviztAPI](https://github.com/Boavizta/boaviztapi) to measure the impact of each service. Boavizta provide their own front-end for the API via [Datavizta](https://dataviz.boavizta.org/serversimpact).

## Details

The architecture is split into three layers:

1. API (+ gateway) - allowing users to query the impact of their usage of cloud services
2. Mapping - maps the usage of cloud services to usage of underlying devices (i.e. servers)
3. Impact - passes the underlying device usage to the BoaviztAPI
