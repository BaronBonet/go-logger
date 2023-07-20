# Go logger

Go logger is a logging port and set of opinionated adapters meant to simplify using a logger in hexagonal architecture.

## Why

In hexagonal architecture external dependencies should be injected into adapters and/or the service at runtime. This means business logic doesn't directly depend on external concerns which can (arguably) make code easier to test and maintain. A logger is an external dependency, therefor it too can be passed into an adapter or the service at runtime. 

```go
import (
	"github.com/BaronBonet/go-logger/logger"
    "github.com/BaronBonet/example-project/internal/core/ports"
)

type service struct {
	logger logger.Logger
	repository ports.Repository
}

func NewService(
	logger logger.Logger
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
