package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	User
	Balance float64 `json:"balance,omitempty"`
}

func (account *Account) PreparUser() error {
	Secret, err := bcrypt.GenerateFromPassword([]byte(account.Secret), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	account.Secret = string(Secret)
	account.CreatedAt = time.Now()

	return nil
}

func NewAccount(name string, cpf string, secret string, balance float64) (*Account, error) {

	account := &Account{
		Balance: balance,
		User:    User{Name: name, CPF: cpf, Secret: secret},
	}

	err := account.PreparUser()

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (account *Account) IsCorrectPassword(secret string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(secret))
	return err == nil
}
