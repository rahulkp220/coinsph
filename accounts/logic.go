package accounts

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	uuid "github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s service) CreateAccount(ctx context.Context, username string, currency string, balance string) (string, error) {
	logger := log.With(s.logger, "method", "CreateAccount")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	account := Account{
		Id:       id,
		Username: username,
		Currency: currency,
		Balance:  balance,
	}

	fmt.Println(account)
	if err := s.repository.CreateAccount(ctx, account); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create account", id)
	return "success", nil
}

func (s service) GetAccount(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetAccount")

	username, err := s.repository.GetAccount(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("get account", id)
	return username, nil
}

func (s service) GetAccounts(ctx context.Context) ([]Account, error) {
	logger := log.With(s.logger, "method", "GetAccount")

	result, err := s.repository.GetAccounts(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return []Account{}, err
	}

	logger.Log("get accounts")
	return result, nil
}

func (s service) DeleteAccount(ctx context.Context, id string) (string, error) {
	userid := id
	logger := log.With(s.logger, "method", "DeleteAccount")

	if err := s.repository.DeleteAccount(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("delete account", userid)
	return "success", nil
}

func (s service) Transfer(ctx context.Context, sender string, reciever string, amount float64) (string, error) {
	logger := log.With(s.logger, "method", "Transfer")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	payment := Payment{
		Id:        id,
		Sender:    sender,
		Reciever:  reciever,
		Amount:    amount,
		Initiated: time.Now(),
	}

	if err := s.repository.Transfer(ctx, payment); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("transfer")
	return "success", nil
}

func (s service) GetTransfers(ctx context.Context) ([]Payment, error) {
	logger := log.With(s.logger, "method", "GetTransfers")

	result, err := s.repository.GetTransfers(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return []Payment{}, err
	}

	logger.Log("get transfers")
	return result, nil
}
