package entity

import "time"

type Bank string

func (b Bank) String() string {
	return string(b)
}

const (
	BankVCB Bank = "VCB"
	BankACB Bank = "ACB"
	BankVIB Bank = "VIB"
)

var SupportedBanks = []string{
	BankVCB.String(),
	BankACB.String(),
	BankVIB.String(),
}

type Account struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Bank      Bank      `json:"bank"`
	UserID    int64     `json:"user_id" db:"user_id"`
	User      *User     `json:"user" db:"user"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}
