# Development

## Requirements

- Install everything you need for a Go [gRPC](https://grpc.io/docs/languages/go/quickstart/) project
- Install [`grpc-gateway`](https://github.com/grpc-ecosystem/grpc-gateway)

## Running locally in Docker

```
docker compose up -d
```

Then go to http://localhost:8081 in your browser.

## Boavizta backend

You can run the [BoaviztAPI](https://github.com/Boavizta/boaviztapi) as follows:

```bash
make boavizta
```

Once started you should be able to go to http://localhost:5000 in your browser to see the API.

You can look at the API docs to see the specification including JSON model: http://localhost:5000/docs.

## Running outside Docker

### Config file

Set up a config file as a copy of the one at the root:

```bash
mkdir -p ~/.config/scw
cp carbon.yml ~/.config/scw/
```

Then change all hosts to `localhost` in this file.

### Server and gateway

Build and run the server and gateway (separate terminals):

```bash
make server
```

Run the gateway:

```bash
make gateway
```

Then access the UI by opening the HTML file, e.g.

```bash
xdg-open site/index.html
```

### Tests

Run the tests and integration tests:

```bash
make test
make test-int
```

## Deploying

The deployment will copy your local checkout to a VM, build containers, then start the application.

Run this with:

```bash
make vm-deploy
```
