package action

import (
	"context"
	"encoding/json"
	"log/slog"
	"strconv"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type action struct {
	con *pgxpool.Pool
}

func NewAction(con *pgxpool.Pool) action {
	return action{con}
}

func (a action) InsertTransaction(body []byte, ctx context.Context) error {
	var (
		data struct {
			Transaction entity.Transaction
			NewBalance  int
		}
	)
	json.Unmarshal(body, &data)

	transaction := data.Transaction
	newBalance := data.NewBalance

	con, err := a.con.Acquire(ctx)
	defer con.Release()

	if err != nil {
		return err
	}

	tx, errTx := con.Begin(ctx)

	if errTx != nil {
		return errTx
	}

	_, errTransacion := tx.Exec(
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

	_, errClient := tx.Exec(
		ctx,
		"UPDATE client SET client_account_balance = $1 WHERE client_id = $2",
		newBalance,
		transaction.TransactionClientID,
	)

	if errTransacion != nil {
		tx.Rollback(ctx)
		return errTransacion
	}

	if errClient != nil {
		tx.Rollback(ctx)
		return errClient
	}

	tx.Commit(ctx)

	slog.Info(
		"[InsertTransaction] " +
			"Transaction with value: " + strconv.Itoa(transaction.TransactionValue) +
			" from Client " + strconv.Itoa(transaction.TransactionClientID) +
			" Saved on Database.",
	)
	return nil
}
