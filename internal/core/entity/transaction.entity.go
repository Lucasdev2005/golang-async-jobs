package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	credit = "c"
	debit  = "d"
)

var vld = validator.New()

func init() {
	vld.RegisterValidation("validTypeTransaction", ValidTypeOfTransaction)
}

type Transaction struct {
	TransactionID          int        `json:"transaction_id"`
	TransactionCreatedAt   time.Time  `json:"transaction_createdAt"`
	TransactionUpdatedAt   time.Time  `json:"transaction_updatedAt"`
	TransactionDeletedAt   *time.Time `json:"transaction_deletedAt"`
	TransactionValue       int        `json:"transaction_value" validate:"required"`
	TransactionType        string     `json:"transaction_type" validate:"required,validTypeTransaction,len=1"`
	TransactionDescription string     `json:"transaction_description" validate:"required,min=10"`
	TransactionClientID    int        `json:"transaction_client_id" validate:"required"`
}

func (t Transaction) IsCredit() bool {
	return t.TransactionType == credit
}

func (t Transaction) IsDebit() bool {
	return t.TransactionType == debit
}

func (t Transaction) ValidTransaction() (bool, error) {
	result := vld.Struct(t)

	if result == nil {
		return true, nil
	} else {
		return false, result
	}
}

func ValidTypeOfTransaction(fl validator.FieldLevel) bool {
	expectedValues := map[string]bool{"c": true, "d": true}
	value := fl.Field().String()
	_, ok := expectedValues[value]
	return ok
}
