package repository

import (
	"context"
	"fmt"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RabbitMQ interface {
	PublishTransaction(transaction entity.Transaction, newBalance int)
	ConsumeMessages(worker func(body []byte, context context.Context) error)
}

type TransactionRepository struct {
	db       *pgxpool.Pool
	rabbitMq RabbitMQ
}

func NewTransactionRepository(db *pgxpool.Pool, rabbitMq RabbitMQ) TransactionRepository {
	return TransactionRepository{
		db:       db,
		rabbitMq: rabbitMq,
	}
}

func (t TransactionRepository) CreateTransaction(transaction entity.Transaction) error {
	var (
		client entity.Client
	)
	t.db.QueryRow(
		context.Background(),
		`
			SELECT 
				client_id, 
				client_account_limit, 
				client_account_balance
			FROM client 
			WHERE client_id = $1
		`,
		transaction.TransactionClientID,
	).Scan(
		&client.ClientID,
		&client.AccountLimit,
		&client.AccountBalance,
	)

	if client.Exists() {
		if ok, errorFromValidatorTransaction := transaction.ValidTransaction(); !ok {
			return errorFromValidatorTransaction
		}

		if balance, ok := client.HaveLimitForTransaction(transaction.TransactionType, transaction.TransactionValue); ok {
			client.AccountBalance = balance
			defer t.rabbitMq.PublishTransaction(transaction, client.AccountBalance)
			return nil
		} else {
			return fmt.Errorf("user no have balance from this transaction")
		}
	} else {
		return fmt.Errorf("client not exists")
	}
}
