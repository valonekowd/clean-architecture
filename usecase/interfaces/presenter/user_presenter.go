package presenter

import (
	"context"

	"github.com/valonekowd/clean-architecture/domain/entity"
	userResp "github.com/valonekowd/clean-architecture/usecase/response/user"
)

type Presenter struct {
	User UserPresenter
}

type UserPresenter interface {
	FetchTransactions(context.Context, []*entity.Transaction) *userResp.GetTransactions
	CreateTransaction(context.Context, *entity.Transaction) *userResp.CreateTransaction
	SignUp(context.Context, *entity.User) (*userResp.SignUp, error)
	SignIn(context.Context, *entity.User) (*userResp.SignIn, error)
}
