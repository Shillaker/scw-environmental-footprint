# Development

## Requirements

- Install everything you need for a Go [gRPC](https://grpc.io/docs/languages/go/quickstart/) project
- Install [`grpc-gateway`](https://github.com/grpc-ecosystem/grpc-gateway)

## Running locally in Docker

```bash
task dev-up
```

Then go to http://localhost:80 in your browser (note HTTP not HTTPS).

## Logging

Logging is managed with Logrus, and the level can be set with the `SCW_IMPACT_LOGGER_LEVEL` environment variable (e.g. to `debug`/`trace`).

## Running outside Docker

### Use local builds

This is easiest using [direnv](https://direnv.net/). Create a `.envrc` file in the root of the project, and add:

```bash
export SCW_IMPACT_GATEWAY_BACKEND_HOST=localhost
export SCW_IMPACT_BOAVIZTA_HOST=localhost

export SCW_IMPACT_GLOBAL_PROJECT_ROOT=<path to this checkout>
```

Then run `direnv allow`.

### Boavizta backend

You can run the [BoaviztAPI](https://github.com/Boavizta/boaviztapi) as follows:

```bash
task boavizta
```

Once started you should be able to go to http://localhost:5000 in your browser to see the API.

You can look at the API docs to see the specification including JSON model: http://localhost:5000/docs.

### Server and gateway

Build and run the server and gateway (separate terminals):

```bash
task server
```

Run the gateway:

```bash
task gateway
```

Run the NGINX proxy and boavizta in the background

```bash
task dev-nginx boavizta
```

Then access the UI by opening the HTML file, e.g.

```bash
xdg-open site/index.html
```

### Tests

Run the tests and integration tests:

```bash
task test
task test-e2e
```

## Deploying

The deployment will copy your local checkout to a VM, build containers, then start the application.

Run this with:

```bash
task vm-deploy
```
