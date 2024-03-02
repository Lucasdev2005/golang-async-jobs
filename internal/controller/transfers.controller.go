package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/Lucasdev2005/golang-async-jobs/internal/cache"
	"github.com/Lucasdev2005/golang-async-jobs/internal/database"
	"github.com/Lucasdev2005/golang-async-jobs/internal/enums"
	"github.com/Lucasdev2005/golang-async-jobs/internal/types"
)

var usersCache = cache.InitCache()

func CreateTransfer(request types.Request) (interface{}, *types.Error) {

	var (
		userId      = request.GetParam("id")
		user        types.User
		transaction types.Transaction
	)

	request.Body(&transaction)
	userFromCache, existsOnCache := usersCache.Get(userId)
	if existsOnCache {
		if u, ok := userFromCache.(types.User); ok {
			user = u
		}
	} else {
		database.Connection.QueryRow(context.Background(), `
			SELECT user_id, user_account_limit, user_account_balance
			FROM "user" WHERE user_id = $1`,
			userId,
		).Scan(&user.UserID, &user.AccountLimit, &user.AccountBalance)
	}

	if user.Exists() {
		if ok, errorFromValidatorTransaction := transaction.ValidTransaction(); !ok {
			return nil, &types.Error{
				ErrorCode: enums.BadRequest,
				Message:   errorFromValidatorTransaction.Error(),
			}
		}

		if balance, ok := user.HaveLimitForTransaction(transaction.TransactionType, transaction.TransactionValue); ok {
			user.AccountBalance = balance
			usersCache.Set(strconv.Itoa(user.UserID), user, 5*time.Minute)
			return user, nil
		}
	}

	// message := types.Message{Content: "Hello, RabbitMQ!"}
	// body, _ := json.Marshal(message)

	// rabbitMq.TransfersChannel.Publish(
	// 	"",
	// 	rabbitMq.TransfersQeue.Name,
	// 	false,
	// 	false,
	// 	amqp.Publishing{
	// 		ContentType: "application/json",
	// 		Body:        body,
	// 	})
	return nil, nil
}
