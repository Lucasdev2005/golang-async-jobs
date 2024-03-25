package actions

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/database"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/entity"
)

func InsertTransaction(body []byte, ctx context.Context) error {
	var (
		data struct {
			Transaction entity.Transaction
			NewBalance  int
		}
	)
	json.Unmarshal(body, &data)

	transaction := data.Transaction
	newBalance := data.NewBalance

	database.Connect()
	defer database.Close()

	_, errTransacion := database.Connection.Exec(
		ctx,
		`INSERT INTO transaction (
			transaction_value,
			transaction_type,
			transaction_description,
			transaction_client_id
		)
		VALUES ($1, $2, $3, $4)`,
		transaction.TransactionValue,
		transaction.TransactionType,
		transaction.TransactionDescription,
		transaction.TransactionClientID,
	)

	database.Connection.Exec(
		ctx,
		"UPDATE client SET client_account_balance = $1 WHERE client_id = $2",
		newBalance,
		transaction.TransactionClientID,
	)

	fmt.Println(
		"[InsertTransaction]",
		"Transaction with value", transaction.TransactionValue,
		"from Client ", transaction.TransactionClientID,
		"Saved on Database.",
	)

	return errTransacion
}
