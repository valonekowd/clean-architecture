package transaction

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/valonekowd/clean-architecture/domain/entity"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/gateway"
)

type sqlRepository struct {
	db   *sqlx.DB
	next gateway.TransactionDataSource
}

var _ gateway.TransactionDataSource = sqlRepository{}

func NewSQLRepository(db *sqlx.DB, next gateway.TransactionDataSource) gateway.TransactionDataSource {
	return sqlRepository{
		db:   db,
		next: next,
	}
}

func (tp sqlRepository) inUse() bool {
	return tp.db != nil
}

func (tp sqlRepository) List(ctx context.Context, userID, accountID int64) ([]*entity.Transaction, error) {
	if !tp.inUse() {
		return tp.next.List(ctx, userID, accountID)
	}

	query := `
		SELECT
			a.id AS "account.id",
			a.bank AS "account.bank",
			t.id,
			t.amount,
			t.type,
			t.created_at
		FROM
			accounts a
			JOIN transactions t ON t.account_id = a.id
		WHERE
			a.user_id = :user_id
	`

	if accountID > 0 {
		query = fmt.Sprintf("%s %s", query, "AND a.id = :account_id")
	}

	arg := map[string]interface{}{
		"user_id":    userID,
		"account_id": accountID,
	}

	fmt.Println(arg)

	stmt, err := tp.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	var ts []*entity.Transaction
	err = stmt.SelectContext(ctx, &ts, arg)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (tp sqlRepository) Add(ctx context.Context, t *entity.Transaction) error {
	if !tp.inUse() {
		return tp.next.Add(ctx, t)
	}

	return nil
}
