package usecase

import (
	"github.com/valonekowd/clean-architecture/infrastructure/log"
	"github.com/valonekowd/clean-architecture/usecase/interactor"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/gateway"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/presenter"
)

type Usecase struct {
	User interactor.UserInteractor
}

func New(ds gateway.DataSource, presenter presenter.Presenter, logger log.Logger) Usecase {
	return Usecase{
		User: interactor.NewUserInteractor(ds, presenter.User, logger),
	}
}
