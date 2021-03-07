package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/valonekowd/clean-architecture/infrastructure/auth"
	"github.com/valonekowd/clean-architecture/infrastructure/log"
	"github.com/valonekowd/clean-architecture/usecase"
)

type ServerEndpoint struct {
	HealthCheck endpoint.Endpoint
	User        UserEndpoint
}

func MakeServerEndpoint(u usecase.Usecase, authCfg auth.Config, logger log.Logger) ServerEndpoint {
	return ServerEndpoint{
		HealthCheck: MakeHealthCheckEndpoint(),
		User:        NewUserEndpoint(u, authCfg, logger),
	}
}

func MakeHealthCheckEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return map[string]string{
			"name":   "E-Banking",
			"status": "SERVING",
		}, nil
	}
}
