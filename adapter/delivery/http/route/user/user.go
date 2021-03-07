package user

import (
	httpTransport "github.com/go-kit/kit/transport/http"

	userJsonDecoder "github.com/valonekowd/clean-architecture/adapter/delivery/http/decode/json/user"
	jsonEncoder "github.com/valonekowd/clean-architecture/adapter/delivery/http/encode/json"
	"github.com/valonekowd/clean-architecture/adapter/endpoint"
	"github.com/valonekowd/clean-architecture/infrastructure/router"
	"github.com/valonekowd/clean-architecture/infrastructure/validator"
)

func New(e endpoint.ServerEndpoint, options []httpTransport.ServerOption, v validator.Validator) func(router.Router) {
	return func(r router.Router) {
		r.Get("/{userID}/transactions", httpTransport.NewServer(
			e.User.GetTransactions,
			userJsonDecoder.ValidatingMiddleware(userJsonDecoder.GetTransactions, v.Struct),
			jsonEncoder.EncodeResponse,
			options...,
		).ServeHTTP)

		r.Post("/{userID}/transactions", httpTransport.NewServer(
			e.User.CreateTransaction,
			userJsonDecoder.ValidatingMiddleware(userJsonDecoder.CreateTransaction, v.Struct),
			jsonEncoder.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
