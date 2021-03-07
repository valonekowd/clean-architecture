package endpoint

import (
	"context"
	"time"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"

	"github.com/valonekowd/clean-architecture/infrastructure/auth"
	"github.com/valonekowd/clean-architecture/infrastructure/log"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())

			return next(ctx, req)
		}
	}
}

func JWTAuthMiddleware(authCfg auth.Config) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		m := jwt.NewParser(authCfg.JWT.Keyfunc, authCfg.JWT.SigningMethod, jwt.StandardClaimsFactory)(next)

		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return m(ctx, req)
		}
	}
}
