package main

import (
	"github.com/Lucasdev2005/golang-async-jobs/internal/consumer/actions"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/rabbitMq"
)

func main() {
	rabbitMq.ConnectionRabbitMq()
	rabbitMq.InitTransfers()
	rabbitMq.ConsumeMessages(actions.InsertTransaction)
}
