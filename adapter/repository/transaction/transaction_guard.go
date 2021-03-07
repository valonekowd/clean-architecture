package transaction

import (
	"context"

	"github.com/valonekowd/clean-architecture/domain/entity"
	"github.com/valonekowd/clean-architecture/domain/errors"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/gateway"
)

type GuardRepository struct{}

var _ gateway.TransactionDataSource = (*GuardRepository)(nil)

func (tp *GuardRepository) List(ctx context.Context, userID, accountID int64) ([]*entity.Transaction, error) {
	return nil, errors.ErrNoDataSource
}

func (tp *GuardRepository) Add(ctx context.Context, t *entity.Transaction) error {
	return errors.ErrNoDataSource
}
