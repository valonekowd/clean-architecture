package user

import (
	"github.com/go-kit/kit/endpoint"
)

var (
	_ endpoint.Failer = GetTransactions{}
	_ endpoint.Failer = CreateTransaction{}
)

type TransactionData struct {
	ID              int64   `json:"id"`
	AccountID       int64   `json:"account_id"`
	Amount          float64 `json:"amount"`
	Bank            string  `json:"bank"`
	TransactionType string  `json:"transaction_type"`
	CreatedAt       string  `json:"created_at"`
}

type GetTransactions struct {
	Data []*TransactionData `json:"data"`
	Err  error              `json:"-"`
}

func (r GetTransactions) Failed() error {
	return r.Err
}

type CreateTransaction struct {
	Data *TransactionData `json:"data"`
	Err  error            `json:"-"`
}

func (r CreateTransaction) Failed() error {
	return r.Err
}

type SignUpPayload struct {
	ID          int64
	Email       string
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Gender      string
	DateOfBirth string `json:"date_of_birth"`
	CreatedAt   string `json:"created_at"`
	AccessToken string `json:"access_token"`
}

type SignUp struct {
	Data *SignUpPayload `json:"data"`
	Err  error          `json:"-"`
}

func (r SignUp) Failed() error {
	return r.Err
}

type SignInPayload struct {
	AccessToken string `json:"access_token"`
}

type SignIn struct {
	Data *SignInPayload `json:"data"`
	Err  error          `json:"-"`
}

func (r SignIn) Failed() error {
	return r.Err
}
