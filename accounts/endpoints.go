package accounts

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateAccount endpoint.Endpoint
	GetAccount    endpoint.Endpoint
	GetAccounts   endpoint.Endpoint
	DeleteAccount endpoint.Endpoint
	Transfer      endpoint.Endpoint
	GetTransfers  endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateAccount: makeCreateAccountEndpoint(s),
		GetAccount:    makeGetAccountEndpoint(s),
		GetAccounts:   makeGetAccountsEndpoint(s),
		DeleteAccount: makeDeleteAccountEdpoint(s),
		Transfer:      makeTransfer(s),
		GetTransfers:  makeGetTransfers(s),
	}
}

func makeCreateAccountEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		ok, err := s.CreateAccount(ctx, req.Username, req.Currency, req.Balance)
		return CreateAccountResponse{Ok: ok}, err
	}
}

func makeGetAccountEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAccountRequest)
		username, err := s.GetAccount(ctx, req.Id)
		return GetAccountResponse{
			Username: username,
		}, err
	}
}

func makeGetAccountsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetAccounts(ctx)
		return GetAccountsResponse{Result: result}, err
	}
}

func makeDeleteAccountEdpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAccountRequest)
		ok, err := s.DeleteAccount(ctx, req.Id)
		return DeleteAccountResponse{Ok: ok}, err
	}
}

func makeTransfer(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TransferRequest)
		ok, err := s.Transfer(ctx, req.Sender, req.Reciever, req.Amount)
		return TransferResponse{Ok: ok}, err
	}
}

func makeGetTransfers(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetTransfers(ctx)
		return GetTransfersResponse{Result: result}, err
	}
}
