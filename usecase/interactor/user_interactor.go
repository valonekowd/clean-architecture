package interactor

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/valonekowd/clean-architecture/domain/entity"
	"github.com/valonekowd/clean-architecture/infrastructure/log"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/gateway"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/presenter"
	userReq "github.com/valonekowd/clean-architecture/usecase/request/user"
	userResp "github.com/valonekowd/clean-architecture/usecase/response/user"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailNotFound          = errors.New("email not found")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)

type UserInteractor interface {
	FetchTransactions(context.Context, userReq.GetTransactions) (*userResp.GetTransactions, error)
	CreateTransaction(context.Context, userReq.CreateTransaction) (*userResp.CreateTransaction, error)
	SignUp(context.Context, userReq.SignUp) (*userResp.SignUp, error)
	SignIn(context.Context, userReq.SignIn) (*userResp.SignIn, error)
}

func NewUserInteractor(ds gateway.DataSource, presenter presenter.UserPresenter, logger log.Logger) UserInteractor {
	var u UserInteractor
	{
		u = NewBasicUserInteractor(ds, presenter)
		// svc = LoggingMiddleware(logger)(svc)
	}
	return u
}

type basicUserInteractor struct {
	ds        gateway.DataSource
	presenter presenter.UserPresenter
}

var _ UserInteractor = &basicUserInteractor{}

func NewBasicUserInteractor(ds gateway.DataSource, presenter presenter.UserPresenter) UserInteractor {
	return basicUserInteractor{
		ds:        ds,
		presenter: presenter,
	}
}

func (b basicUserInteractor) FetchTransactions(ctx context.Context, req userReq.GetTransactions) (*userResp.GetTransactions, error) {
	ts, err := b.ds.Transaction.List(ctx, req.UserID, req.AccountID)
	if err != nil {
		return nil, err
	}

	return b.presenter.FetchTransactions(ctx, ts), nil
}

func (b basicUserInteractor) CreateTransaction(ctx context.Context, req userReq.CreateTransaction) (*userResp.CreateTransaction, error) {
	t := &entity.Transaction{
		Type:      entity.TransactionType(req.TransactionType),
		Amount:    req.Amount,
		AccountID: req.AccountID,
	}

	if err := b.ds.Transaction.Add(ctx, t); err != nil {
		return nil, err
	}

	return b.presenter.CreateTransaction(ctx, t), nil
}

func (b basicUserInteractor) SignUp(ctx context.Context, req userReq.SignUp) (*userResp.SignUp, error) {
	u := &entity.User{
		Email:       req.Email,
		Password:    req.Password,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Gender:      entity.Gender(req.Gender),
		DateOfBirth: req.DateOfBirth,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := u.HashPassword(); err != nil {
		return nil, fmt.Errorf("hashing user password: %w", err)
	}

	if err := b.ds.User.Add(ctx, u); err != nil {
		return nil, fmt.Errorf("inserting user: %w", err)
	}

	return b.presenter.SignUp(ctx, u)
}

func (b basicUserInteractor) SignIn(ctx context.Context, req userReq.SignIn) (*userResp.SignIn, error) {
	u, err := b.ds.User.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("finding user by email %s: %w", req.Email, err)
	}

	if u == nil {
		return nil, fmt.Errorf("finding user by email %s: %w", req.Email, ErrEmailNotFound)
	}

	if isMatch, err := u.ComparePassword(req.Password); !isMatch {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, fmt.Errorf("comparing user password: %w", ErrInvalidEmailOrPassword)
		}
		return nil, fmt.Errorf("comparing user password: %w", err)
	}

	return b.presenter.SignIn(ctx, u)
}
