package main

import (
	action "github.com/Lucasdev2005/golang-async-jobs/internal/consumer/actions"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/database"
	rabbitmq "github.com/Lucasdev2005/golang-async-jobs/internal/core/rabbitMq"
)

func main() {
	database := database.NewDatabase()
	rabbitMq := rabbitmq.NewRabbitMq()
	action := action.NewAction(database.Con)

	rabbitMq.ConsumeMessages(action.InsertTransaction)
}
