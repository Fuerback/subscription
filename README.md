# Subscription
The project is a Subscription service using Golang and REST API that uses gorilla/mux as router framework, SQLite as a database and Docker to build.

The project is also using Clean Architecture.

# TODO
- [x] Json validator
- [ ] Context timeout
- [ ] Integration test
- [ ] Cache
- [ ] Error handler
- [ ] Test cover
- [ ] Add audit columns in DB

# Getting started

## Locally

To run project locally you need to export the environment variable `PORT=:8080` or to other port to your choice and. 
You must have installed:

- Golang >=1.16

```shell
make run-local
```

```shell
make tests
```

## Docker

```shell
make docker-image
make docker-run
```

or just

```shell
make docker-up
```

The default PORT to test via Docker is `8080`

to run **tests on Docker**:

```shell
make docker-tests
```

## Documentation

The API documentation [./docs/api/swagger.yaml](https://github.com/Fuerback/subscription/blob/main/docs/api/swagger.yaml)
The database schema [./docs/db/v1.0.0-schema.sql](https://github.com/Fuerback/subscription/blob/main/docs/db/v1.0.0-schema.sql)
The database model [./docs/db/model.png](https://github.com/Fuerback/subscription/blob/main/docs/db/model.png)
