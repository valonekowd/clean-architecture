package entity

import "time"

type TransactionType string

func (t TransactionType) String() string {
	return string(t)
}

const (
	TransactionWithdraw TransactionType = "withdraw"
	TransactionDeposit  TransactionType = "deposit"
)

var SupportedTransactionTypes = []string{
	TransactionWithdraw.String(),
	TransactionDeposit.String(),
}

type Transaction struct {
	ID        int64           `json:"id"`
	Type      TransactionType `json:"type"`
	Amount    float64         `json:"amount"`
	AccountID int64           `json:"account_id" db:"account_id"`
	Account   *Account        `json:"account" db:"account"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}
