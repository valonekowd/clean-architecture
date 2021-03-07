package formatter

import (
	"github.com/valonekowd/clean-architecture/adapter/formatter/user"
	"github.com/valonekowd/clean-architecture/infrastructure/auth"
	"github.com/valonekowd/clean-architecture/infrastructure/log"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/presenter"
)

type Formatter struct{}

func New(authCfg auth.Config, logger log.Logger) presenter.Presenter {
	return presenter.Presenter{
		User: user.NewFormatter(authCfg, logger),
	}
}
