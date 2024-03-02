package types

import (
	"time"
)

type User struct {
	UserID         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
	AccountLimit   int
	AccountBalance int
}

func (u User) Exists() bool {
	return u.UserID != 0
}

func (u User) HaveLimitForTransaction(transactionType string, transactionValue int) (int, bool) {
	var balance int

	if transactionType == debit {
		balance = u.AccountBalance - transactionValue
	}
	if transactionType == credit {
		balance = u.AccountBalance + transactionValue
	}

	return balance, (u.AccountLimit + balance) > 0
}
