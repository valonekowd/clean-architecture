package http

import (
	"net/http"

	"github.com/go-kit/kit/auth/jwt"
	kitTransport "github.com/go-kit/kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"

	jsonEncoder "github.com/valonekowd/clean-architecture/adapter/delivery/http/encode/json"
	"github.com/valonekowd/clean-architecture/adapter/delivery/http/option"
	userRouter "github.com/valonekowd/clean-architecture/adapter/delivery/http/route/user"
	"github.com/valonekowd/clean-architecture/adapter/endpoint"
	"github.com/valonekowd/clean-architecture/adapter/middleware"
	"github.com/valonekowd/clean-architecture/infrastructure/log"
	"github.com/valonekowd/clean-architecture/infrastructure/router"
	"github.com/valonekowd/clean-architecture/infrastructure/router/chi"
	"github.com/valonekowd/clean-architecture/infrastructure/validator"
)

const (
	apiVersion0 = ""
)

func NewHTTPHandler(e endpoint.ServerEndpoint, v validator.Validator, logger log.Logger) http.Handler {
	var r router.Router
	{
		r = chi.NewRouter()
		r.Use(middleware.CORSMiddleware)
	}

	options := []httpTransport.ServerOption{
		httpTransport.ServerBefore(
			option.LogRequestInfo(logger),
			jwt.HTTPToContext(),
		),
		httpTransport.ServerErrorHandler(kitTransport.NewLogErrorHandler(logger)),
		httpTransport.ServerErrorEncoder(jsonEncoder.EncodeError),
	}

	r.Get("/", httpTransport.NewServer(
		e.HealthCheck,
		httpTransport.NopRequestDecoder,
		httpTransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/api", func(r router.Router) {
		r.Route("/"+apiVersion0, func(r router.Router) {
			r.Route("/users", userRouter.New(e, options, v))
		})
	})

	return r
}
