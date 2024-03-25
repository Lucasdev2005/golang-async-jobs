package entity

import (
	"time"
)

type Client struct {
	ClientID       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	AccountLimit   int
	AccountBalance int
}

func (u Client) Exists() bool {
	return u.ClientID != 0
}

func (u Client) HaveLimitForTransaction(transactionType string, transactionValue int) (int, bool) {
	var balance int

	if transactionType == debit {
		balance = u.AccountBalance - transactionValue
	}
	if transactionType == credit {
		balance = u.AccountBalance + transactionValue
	}

	return balance, (u.AccountLimit + balance) > 0
}
