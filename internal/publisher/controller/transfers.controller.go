package controller

import (
	"log"
	"strconv"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/entity"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/enums"
	"github.com/Lucasdev2005/golang-async-jobs/internal/publisher/repository"
)

var transactionrepository = repository.NewTransactionRepository()

func PublishTransfer(request entity.Request) (interface{}, *entity.Error) {

	var (
		clientId    = request.GetParam("id")
		transaction entity.Transaction
	)

	request.Body(&transaction)
	transaction.TransactionClientID, _ = strconv.Atoi(clientId)
	errorOnSaveTransaction := transactionrepository.Create(transaction)

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
