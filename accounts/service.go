package accounts

import (
	"context"
)

// service method exposed
type Service interface {

	// amount CRUD
	CreateAccount(ctx context.Context, username string, currency string, balance string) (string, error)
	GetAccount(ctx context.Context, id string) (string, error)
	GetAccounts(ctx context.Context) ([]Account, error)
	DeleteAccount(ctx context.Context, id string) (string, error)

	// payment CRUD
	Transfer(ctx context.Context, sender string, reciever string, Amount float64) (string, error)
	GetTransfers(ctx context.Context) ([]Payment, error)
}
