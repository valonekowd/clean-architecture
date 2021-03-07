package transaction

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/valonekowd/clean-architecture/domain/entity"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/gateway"
)

type mongoRepository struct {
	db   *mongo.Database
	next gateway.TransactionDataSource
}

var _ gateway.TransactionDataSource = mongoRepository{}

func NewMongoRepository(db *mongo.Database, next gateway.TransactionDataSource) gateway.TransactionDataSource {
	return mongoRepository{
		db:   db,
		next: next,
	}
}

func (tp mongoRepository) inUse() bool {
	return tp.db != nil
}

func (tp mongoRepository) List(ctx context.Context, userID, accountID int64) ([]*entity.Transaction, error) {
	if !tp.inUse() {
		return tp.next.List(ctx, userID, accountID)
	}

	return nil, nil
}

func (tp mongoRepository) Add(ctx context.Context, t *entity.Transaction) error {
	if !tp.inUse() {
		return tp.next.Add(ctx, t)
	}

	return nil
}
