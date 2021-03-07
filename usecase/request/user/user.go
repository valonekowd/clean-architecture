package user

import "time"

type GetTransactions struct {
	UserID    int64 `validate:"required,gt=0"`
	AccountID int64 `validate:"gt=0"`
}

type CreateTransaction struct {
	UserID          int64   `validate:"required,gt=0"`
	AccountID       int64   `validate:"required,gt=0" json:"account_id"`
	Amount          float64 `validate:"required,gt=0" json:"amount"`
	TransactionType string  `validate:"required,is-valid-transaction-type" json:"transaction_type"`
}

type SignUp struct {
	Email       string    `validate:"required,email,max=300"`
	Password    string    `validate:"required,min=8"`
	FirstName   string    `validate:"required,min=8,max=50" json:"first_name"`
	LastName    string    `validate:"required,min=8,max=50" json:"last_name"`
	Gender      string    `validate:"required,is-valid-gender"`
	DateOfBirth time.Time `validate:"required" json:"date_of_birth"`
}

type SignIn struct {
	Email    string `validate:"required,email,max=300"`
	Password string `validate:"required,min=8"`
}
