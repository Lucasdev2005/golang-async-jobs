package rabbitMq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/types"
	"github.com/streadway/amqp"
)

var (
	MqConnection     *amqp.Connection
	TransfersChannel *amqp.Channel
	TransfersQueue   amqp.Queue
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ConnectionRabbitMq() {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	MqConnection = conn
}

func InitTransfers() {
	c, _ := MqConnection.Channel()
	TransfersChannel = c

	TransfersQueue, _ = TransfersChannel.QueueDeclare(
		"transferencia",
		false,
		false,
		false,
		false,
		nil,
	)
}

func PublishTransaction(transaction types.Transaction, newBalance int) {
	data := struct {
		Transaction types.Transaction
		NewBalance  int
	}{
		Transaction: transaction,
		NewBalance:  newBalance,
	}
	body, _ := json.Marshal(data)
	err := TransfersChannel.Publish(
		"",
		TransfersQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	failOnError(err, "Failed to publish a message")
}

func ConsumeMessages(worker func(body []byte)) {
	TransfersChannel.Qos(1, 0, false)
	msgs, err := TransfersChannel.Consume(
		TransfersQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")
	fmt.Println("Waiting messages...")

	forever := make(chan bool)
	func() {
		for d := range msgs {
			go worker(d.Body)
		}
	}()

	<-forever
}
