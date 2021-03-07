package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Gender string

func (g Gender) String() string {
	return string(g)
}

const (
	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"
)

var SupportedGenders = []string{
	Male.String(),
	Female.String(),
	Other.String(),
}

type User struct {
	ID          int64      `json:"id"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    string     `json:"last_name" db:"last_name"`
	Gender      Gender     `json:"gender"`
	DateOfBirth time.Time  `json:"date_of_birth" db:"date_of_birth"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

func (u *User) HashPassword() error {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err == nil {
		u.Password = string(b)
	}
	return err
}

func (u *User) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil, err
}
