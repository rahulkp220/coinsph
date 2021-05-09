package accounts

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/shopspring/decimal"
)

var ErrRepo = errors.New("unable to handle repo request")
var ErrBalanceValidation = errors.New("invalid data type supplied for balance")
var ErrInsufficientBalance = errors.New("insufficient balance")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateAccount(ctx context.Context, account Account) error {
	sql := `
		INSERT INTO accounts (id, username, currency, balance)
		VALUES ($1, $2, $3, $4)
	`

	// validation for balance
	val, valerr := decimal.NewFromString(account.Balance)
	if valerr != nil {
		return ErrBalanceValidation
	}
	balance := val.String()

	if account.Username == "" || account.Currency == "" {
		return ErrRepo
	}

	fmt.Println(account.Username, account.Balance, account.Currency)
	_, err := repo.db.ExecContext(ctx, sql, account.Id, account.Username, account.Currency, balance)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetAccount(ctx context.Context, id string) (string, error) {
	var username string
	err := repo.db.QueryRow("SELECT username FROM accounts WHERE id=$1", id).Scan(&username)
	if err != nil && err != sql.ErrNoRows {
		return "", ErrRepo
	}

	return username, nil
}

func (repo *repo) GetAccounts(ctx context.Context) ([]Account, error) {
	var rowmap []Account
	rows, err := repo.db.Query("SELECT id, username, currency, balance FROM accounts")
	if err != nil {
		return []Account{}, ErrRepo
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var username string
		var currency string
		var balance string
		err = rows.Scan(&id, &username, &currency, &balance)
		if err != nil {
			return []Account{}, err
		}

		rowmap = append(rowmap, Account{Id: id, Username: username, Balance: balance, Currency: currency})
	}

	err = rows.Err()
	if err != nil {
		return []Account{}, err
	}

	fmt.Print(rowmap)
	return rowmap, err
}

func (repo *repo) DeleteAccount(ctx context.Context, id string) error {
	_, err := repo.db.Exec("DELETE FROM accounts WHERE id = $1", id)
	if err != nil {
		return ErrRepo
	}

	return nil
}

func (repo *repo) Transfer(ctx context.Context, payment Payment) error {
	fmt.Println(payment)

	sql := `
		INSERT INTO payments (id, sender, reciever, amount, initiated, completed)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	// get sender
	var senderid string
	var senderbalance float64
	var sendercurrency string
	err := repo.db.QueryRow("SELECT id, balance, currency FROM accounts WHERE id=$1", payment.Sender).Scan(&senderid, &senderbalance, &sendercurrency)
	if err != nil {
		return err
	}

	// get reciever
	var recieverid string
	var recieverbalance float64
	var recievercurrency string
	err = repo.db.QueryRow("SELECT id, balance, currency FROM accounts WHERE id=$1", payment.Reciever).Scan(&recieverid, &recieverbalance, &recievercurrency)
	if err != nil {
		return err
	}

	// make a transfer only if sender has enough balance
	maketransfer := big.NewFloat(payment.Amount).Cmp(big.NewFloat(senderbalance))
	newsenderbalance := senderbalance - payment.Amount
	newreceiverbalance := payment.Amount + recieverbalance

	// Terms of transfer
	// 1. Sender has sufficient balance
	// 2. Account cannot be drained to 0
	// 3. Currency for both accounts is same
	if maketransfer < 0 && newsenderbalance > 0.01 && sendercurrency == recievercurrency {

		// convert balance to string for accounts
		newsenderbalancestr := fmt.Sprintf("%f", newsenderbalance)
		newreceiverbalancestr := fmt.Sprintf("%f", newreceiverbalance)

		// create a new payment in an atomic block
		tx, err := repo.db.BeginTx(ctx, nil)
		if err != nil {
			return ErrRepo
		}

		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE id = $2", newsenderbalancestr, senderid)
		if err != nil {
			if rb := tx.Rollback(); rb != nil {
				return rb
			}
			return err
		}

		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE id = $2", newreceiverbalancestr, recieverid)
		if err != nil {
			if rb := tx.Rollback(); rb != nil {
				return rb
			}
			return err
		}

		_, err = tx.ExecContext(ctx, sql, payment.Id, payment.Sender, payment.Reciever, payment.Amount, payment.Initiated, time.Now())
		if err != nil {
			if rb := tx.Rollback(); rb != nil {
				return rb
			}
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
	} else {
		return ErrInsufficientBalance
	}

	return nil

}

func (repo *repo) GetTransfers(ctx context.Context) ([]Payment, error) {
	var rowmap []Payment
	rows, err := repo.db.Query("SELECT id, sender, reciever, amount, initiated, completed FROM payments")
	if err != nil {
		return []Payment{}, ErrRepo
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var sender string
		var reciever string
		var amount float64
		var initiated time.Time
		var completed time.Time
		err = rows.Scan(&id, &sender, &reciever, &amount, &initiated, &completed)
		if err != nil {
			return []Payment{}, err
		}

		rowmap = append(rowmap, Payment{Id: id, Sender: sender, Reciever: reciever, Amount: amount, Initiated: initiated, Completed: completed})
	}

	fmt.Println(rowmap)
	err = rows.Err()
	if err != nil {
		return []Payment{}, err
	}

	return rowmap, err
}
