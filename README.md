# Go logger

Go logger is a logging port and set of opinionated adapters meant to simplify using a logger in hexagonal architecture.

## Install

```shell
go get BaronBonet/go-logger.v1
```

## Why

In hexagonal architecture dependencies external dependencies should be injected to adapters and the service at runtime. This means business logic doesn't directly depend on external concerns which can (arguably) make code easier to test and maintain. A logger is an external dependency, therefor it to can be passed into a service or adapter at runtime. 

```go
type service struct {
  logger logger
  repository ports.Repository
}

func NewService(
  logger logger
  repository ports.Repository
  ) ports.Service {
  return &service{
    logger: logger,
    repository: repository
    }
}
```

I personally prefer to have 3 distinct logger adapters:

1. A logger that provides structured logging for a service running on AWS (or some other cloud provider), to enable parsing of logs.
2. A colorful logger for CLI tools and local development.
3. A logger that can be passed into a service or adapter during testing.
![Go logger example](https://cdn.ericcbonet.com/go-logger.png)
