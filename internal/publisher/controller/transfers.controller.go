package controller

import (
	"log"
	"strconv"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/entity"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/enums"
)

func NewTransactionCotroller(repository TransactionRepository) TransactionController {
	return TransactionController{
		repository,
	}
}

type TransactionRepository interface {
	CreateTransaction(transaction entity.Transaction) error
}

type TransactionController struct {
	repository TransactionRepository
}

func (t TransactionController) PublishTransfer(request entity.Request) (interface{}, *entity.Error) {

	var (
		clientId    = request.GetParam("id")
		transaction entity.Transaction
	)

	request.Body(&transaction)
	transaction.TransactionClientID, _ = strconv.Atoi(clientId)
	errorOnSaveTransaction := t.repository.CreateTransaction(transaction)

	log.Println("[PublishTransfer] errorOnSaveTransaction: ", errorOnSaveTransaction)
	if errorOnSaveTransaction != nil {
		return nil, &entity.Error{
			ErrorCode: enums.NetworkAuthenticationRequired,
			Message:   errorOnSaveTransaction.Error(),
		}
	} else {
		return "Transfer Created.", &entity.Error{}
	}
}
