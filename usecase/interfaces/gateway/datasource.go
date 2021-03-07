package gateway

import (
	"context"

	"github.com/valonekowd/clean-architecture/domain/entity"
)

type DataSource struct {
	Transaction TransactionDataSource
	User        UserDataSource
}

type UserDataSource interface {
	FindByEmail(_ context.Context, email string) (*entity.User, error)
	Add(context.Context, *entity.User) error
}

type TransactionDataSource interface {
	List(_ context.Context, userID, accountID int64) ([]*entity.Transaction, error)
	Add(context.Context, *entity.Transaction) error
}
