# Data

## Generating source data

The source data is pulled from the [Scaleway API](https://www.scaleway.com/en/developers/api), using the [Scaleway Go SDK](https://www.scaleway.com/en/docs/developer-tools/scaleway-sdk/go-sdk/).

To update the data, you can do the following:

- Create a [Scaleway Console](https://console.scaleway.com) account, and an associated API key
- Install [direnv](https://direnv.net/)
- Add a `.envrc` in the root of this directory containing
  - `SCW_IMPACT_SCW_ACCESS_KEY` - the access key for your API key
  - `SCW_IMPACT_SCW_SECRET_KEY` - the secret key for your API key
- Run `direnv allow`
- Run `task source-data`
