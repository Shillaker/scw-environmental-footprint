# ResilioDB

[ResilioDB](https://db.resilio.tech) is a proprietary database of impact data for digital equipment.

You can see some sample requests and responses for their API [here](https://db.resilio.tech/try).

## Setup

To use ResilioDB as a backend, you need to set the backend to `resilio` in your requests, and set the following environment variables (e.g. in your `.envrc`):

```bash
export CARBON_RESILIO_TOKEN=<valid resilio token>
```

