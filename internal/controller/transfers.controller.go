package controller

import (
	"encoding/json"

	"github.com/Lucasdev2005/golang-async-jobs/internal/rabbitMq"
	"github.com/Lucasdev2005/golang-async-jobs/internal/types"
	"github.com/streadway/amqp"
)

func CreateTransfer(request types.Request) (interface{}, *types.Error) {
	message := types.Message{Content: "Hello, RabbitMQ!"}
	body, _ := json.Marshal(message)
	rabbitMq.TransfersChannel.Publish(
		"",
		rabbitMq.TransfersQeue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return nil, nil
}
