package repository

import (
	"context"
	"fmt"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/database"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/entity"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/rabbitMq"
)

type transactionRepository struct {
	Create func(t entity.Transaction) error
}

func NewTransactionRepository() transactionRepository {
	return transactionRepository{
		Create: createTransaction,
	}
}

func createTransaction(transaction entity.Transaction) error {
	var (
		client entity.Client
	)
	database.Connection.QueryRow(
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
			defer rabbitMq.PublishTransaction(transaction, client.AccountBalance)
			return nil
		} else {
			return fmt.Errorf("user no have balance from this transaction")
		}
	} else {
		return fmt.Errorf("client not exists")
	}
}
