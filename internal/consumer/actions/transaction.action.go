package actions

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/database"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/types"
)

func InsertTransaction(body []byte) {
	var (
		data struct {
			Transaction types.Transaction
			NewBalance  int
		}
	)
	json.Unmarshal(body, &data)

	transaction := data.Transaction
	newBalance := data.NewBalance

	database.Connect()

	_, errTransacion := database.Connection.Query(
		context.Background(),
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

	errUpdateBalance := database.Connection.QueryRow(
		context.Background(),
		"UPDATE client SET client_account_balance = $1 WHERE client_id = $2",
		newBalance,
		transaction.TransactionClientID,
	).Scan()

	fmt.Println(
		"[InsertTransaction]",
		"Transaction with value", transaction.TransactionValue,
		"from Client ", transaction.TransactionClientID,
		"Saved on Database.",
	)

	database.Close()
	resolveErr(errTransacion)
	resolveErr(errUpdateBalance)

}

func resolveErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
