# Boavizta API

The Boavizta API uses publicly available data to estimate the impact of devices during two phases:

- Manufacture - impact of building the device
- Use - impact of using the device

It provides three impacts:

- Global warming potential (GWP) (kgCO2E) - heat absorbed in the atmosphere equivalent to a weight of CO2 emitted ([wiki](https://en.wikipedia.org/wiki/Global_warming_potential))
- Primary energy (PE) (MJ) - equivalent energy consumed from soureces found naturally in nature, not yet converted by any human process ([wiki](https://en.wikipedia.org/wiki/Primary_energy))
- Abiotic depletion potential (ADP) (KgSbeq) - measure of environmental impact of extracting materials vs. their global stocks, includes both raw materials and fossil fuels ([link](https://www.sciencedirect.com/topics/engineering/abiotic-depletion-potential))

## Methodology

_Inputs_

1. Server/device configuration - CPU, RAM, SSD etc.
2. Usage - region, time, load

These are sent to the Boavizta `server` API. A visualisation of this with configurable inputs can be found [here](https://dataviz.boavizta.org/serversimpact).

_Calculation_

The inputs to the methodology are:

- **Devices** - set of devices, such as CPUs and SSDs, mapped from the input server configuration (taking default values where not specified)
- **Load** - the load exerted on the server over time, either as a single constant percentage, or a more granular list of loads at different percentages for different times
- **Consumption profile** - how a component's power consumption varies under load, calculated using the user's specified power consumption value(s), or using hard-coded estimations for each component
- **Energy mix** - publicly available data on the energy mix for the region in which the server is used
- **Lifespan** - the lifespan of the underlying hardware - this is used to divide the manufacture cost proportionally between users

These inputs are used to calculate:

- Two impact values (manufacture and use)
- ... for each of the three impact criteria (GWP, PE and ADP)
- ... for each component of the device

I.e. for the CPU, the API will calculate the GWP, PE and ADP for its manufacture and use, same for the RAM and so on.

Finally, the impact of all the components is combined to produce the final impact of the server usage.

## Links

### Docs

Useful links from the [BoaviztAPI docs](https://doc.api.boavizta.org/):

- [Usage methodology](https://doc.api.boavizta.org/Explanations/usage/usage/)
- [Manufacture methodology](https://doc.api.boavizta.org/Explanations/manufacture_methodology/)
- [Electrical impact factors](https://doc.api.boavizta.org/Explanations/usage/elec_factors/)
- [Impact criteria](https://doc.api.boavizta.org/Explanations/impacts/)
- [Components](https://doc.api.boavizta.org/Explanations/components/component/)

### Data

Source data:

- [Electricity data per region](https://github.com/Boavizta/boaviztapi/blob/main/boaviztapi/data/electricity/electricity_impact_factors.csv)
- [Components](https://github.com/Boavizta/boaviztapi/tree/main/boaviztapi/data/components)

### Code

Useful links into the implementation:

- [Consumption profile model](https://github.com/Boavizta/boaviztapi/blob/main/boaviztapi/model/consumption_profile/consumption_profile.py)
- [Component impact calculations](https://github.com/Boavizta/boaviztapi/tree/main/boaviztapi/model/component)
    - [CPU impact](https://github.com/Boavizta/boaviztapi/blob/main/boaviztapi/model/component/cpu.py)
    - [RAM impact](https://github.com/Boavizta/boaviztapi/blob/main/boaviztapi/model/component/ram.py)
    - [SSD impact](https://github.com/Boavizta/boaviztapi/blob/main/boaviztapi/model/component/ssd.py)
