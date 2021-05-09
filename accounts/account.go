package accounts

import (
	"context"
	"time"
)

type Account struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}

type Payment struct {
	Id        string    `json:"id"`
	Sender    string    `json:"sender"`
	Reciever  string    `json:"reciever"`
	Amount    float64   `json:"amount"`
	Initiated time.Time `json:"initiated"`
	Completed time.Time `json:"completed"`
}

// method exposed to database
type Repository interface {

	// account CRUD
	CreateAccount(ctx context.Context, account Account) error
	GetAccount(ctx context.Context, id string) (string, error)
	GetAccounts(ctx context.Context) ([]Account, error)
	DeleteAccount(ctx context.Context, id string) error

	// payment CRUD
	Transfer(ctx context.Context, payment Payment) error
	GetTransfers(ctx context.Context) ([]Payment, error)
}
