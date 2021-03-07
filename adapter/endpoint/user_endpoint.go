package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/valonekowd/clean-architecture/infrastructure/auth"
	"github.com/valonekowd/clean-architecture/infrastructure/log"
	"github.com/valonekowd/clean-architecture/usecase"
	userReq "github.com/valonekowd/clean-architecture/usecase/request/user"
)

type UserEndpoint struct {
	GetTransactions   endpoint.Endpoint
	CreateTransaction endpoint.Endpoint
}

func NewUserEndpoint(u usecase.Usecase, authCfg auth.Config, logger log.Logger) UserEndpoint {
	var getTransactionsEndpoint endpoint.Endpoint
	{
		getTransactionsEndpoint = MakeGetTransactionsEndpoint(u)
		getTransactionsEndpoint = JWTAuthMiddleware(authCfg)(getTransactionsEndpoint)
		getTransactionsEndpoint = LoggingMiddleware(logger)(getTransactionsEndpoint)
	}

	var createTransactionEndpoint endpoint.Endpoint
	{
		createTransactionEndpoint = MakeCreateTransactionEndpoint(u)
		createTransactionEndpoint = JWTAuthMiddleware(authCfg)(createTransactionEndpoint)
		createTransactionEndpoint = LoggingMiddleware(logger)(createTransactionEndpoint)
	}

	return UserEndpoint{
		GetTransactions:   getTransactionsEndpoint,
		CreateTransaction: createTransactionEndpoint,
	}
}

func MakeGetTransactionsEndpoint(u usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(userReq.GetTransactions)

		return u.User.FetchTransactions(ctx, r)
	}
}

func MakeCreateTransactionEndpoint(u usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(userReq.CreateTransaction)

		return u.User.CreateTransaction(ctx, r)
	}
}
