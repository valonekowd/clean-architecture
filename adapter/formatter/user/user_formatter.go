package user

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/valonekowd/clean-architecture/domain/entity"
	"github.com/valonekowd/clean-architecture/infrastructure/auth"
	"github.com/valonekowd/clean-architecture/infrastructure/log"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/presenter"
	userResp "github.com/valonekowd/clean-architecture/usecase/response/user"
)

type userFormatter struct {
	logger  log.Logger
	authCfg auth.Config
}

var _ presenter.UserPresenter = userFormatter{}

func NewFormatter(authCfg auth.Config, logger log.Logger) presenter.UserPresenter {
	return userFormatter{
		logger:  logger,
		authCfg: authCfg,
	}
}

func (f userFormatter) mappingTransactionData(ctx context.Context, t *entity.Transaction) *userResp.TransactionData {
	return &userResp.TransactionData{
		ID:              t.ID,
		AccountID:       t.Account.ID,
		Amount:          t.Amount,
		Bank:            t.Account.Bank.String(),
		TransactionType: t.Type.String(),
		CreatedAt:       t.CreatedAt.Format("2020-02-10 20:00:00 +0700"),
	}
}

func (f userFormatter) FetchTransactions(ctx context.Context, ts []*entity.Transaction) *userResp.GetTransactions {
	data := make([]*userResp.TransactionData, 0, len(ts))

	for _, t := range ts {
		data = append(data, f.mappingTransactionData(ctx, t))
	}

	return &userResp.GetTransactions{Data: data}
}

func (f userFormatter) CreateTransaction(ctx context.Context, t *entity.Transaction) *userResp.CreateTransaction {
	return &userResp.CreateTransaction{Data: f.mappingTransactionData(ctx, t)}
}

func (f userFormatter) getAccessToken(ctx context.Context, u *entity.User) (string, error) {
	token := jwt.NewWithClaims(f.authCfg.JWT.SigningMethod, jwt.MapClaims{
		"iss": f.authCfg.JWT.Issuer,
		"exp": time.Now().Add(time.Hour * 8).Unix(),
		"uid": u.ID,
	})

	key, err := f.authCfg.JWT.Keyfunc(token)
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

func (f userFormatter) SignUp(ctx context.Context, u *entity.User) (*userResp.SignUp, error) {
	token, err := f.getAccessToken(ctx, u)
	if err != nil {
		return nil, err
	}

	return &userResp.SignUp{
		Data: &userResp.SignUpPayload{
			ID:          u.ID,
			CreatedAt:   u.CreatedAt.Format("2020-02-10 20:00:00 +0700"),
			AccessToken: token,
		},
	}, nil
}

func (f userFormatter) SignIn(ctx context.Context, u *entity.User) (*userResp.SignIn, error) {
	token, err := f.getAccessToken(ctx, u)
	if err != nil {
		return nil, err
	}

	return &userResp.SignIn{
		Data: &userResp.SignInPayload{
			AccessToken: token,
		},
	}, nil
}
