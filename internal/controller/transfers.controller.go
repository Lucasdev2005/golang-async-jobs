package controller

import (
	"context"

	"github.com/Lucasdev2005/golang-async-jobs/internal/database"
	"github.com/Lucasdev2005/golang-async-jobs/internal/types"
)

func CreateTransfer(request types.Request) (interface{}, *types.Error) {

	var (
		userId = request.GetParam("id")
		user   types.User
	)

	database.Connection.QueryRow(context.Background(), `
		SELECT user_id, user_account_limit, user_account_balance 
		FROM user WHERE user_id = $1`,
		userId,
	).Scan(&user.UserID, &user.AccountLimit, &user.AccountBalance)

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
	return user, nil
}
