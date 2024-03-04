package controller

import (
	"context"
	"strconv"

	Cache "github.com/Lucasdev2005/golang-async-jobs/internal/core/cache"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/database"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/enums"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/rabbitMq"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/types"
	"github.com/patrickmn/go-cache"
)

func CreateTransfer(request types.Request) (interface{}, *types.Error) {

	var (
		clientId    = request.GetParam("id")
		client      types.Client
		transaction types.Transaction
	)

	request.Body(&transaction)
	transaction.TransactionClientID, _ = strconv.Atoi(clientId)
	clientFromCache, clientExistsOnCache := Cache.ClientsCache.Get(clientId)
	if clientExistsOnCache {
		if u, ok := clientFromCache.(types.Client); ok {
			client = u
		}
	} else {
		database.Connection.QueryRow(context.Background(), `
			SELECT client_id, client_account_limit, client_account_balance
			FROM client WHERE client_id = $1`,
			clientId,
		).Scan(&client.ClientID, &client.AccountLimit, &client.AccountBalance)
	}

	if client.Exists() {
		if ok, errorFromValidatorTransaction := transaction.ValidTransaction(); !ok {
			return nil, &types.Error{
				ErrorCode: enums.BadRequest,
				Message:   errorFromValidatorTransaction.Error(),
			}
		}

		if balance, ok := client.HaveLimitForTransaction(transaction.TransactionType, transaction.TransactionValue); ok {
			client.AccountBalance = balance
			Cache.ClientsCache.Set(clientId, client, cache.DefaultExpiration)
			defer rabbitMq.PublishTransaction(transaction, client.AccountBalance)
			return client, nil
		} else {
			return nil, &types.Error{
				ErrorCode: enums.UnprocessableEntity,
				Message:   "User don't have balance from this transaction.",
			}
		}
	}

	return nil, &types.Error{
		ErrorCode: enums.NetworkAuthenticationRequired,
		Message:   "User not found",
	}
}
