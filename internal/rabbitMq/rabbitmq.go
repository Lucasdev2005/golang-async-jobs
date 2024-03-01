package rabbitMq

import (
	"log"

	"github.com/streadway/amqp"
)

var (
	MqConnection     *amqp.Connection
	TransfersChannel *amqp.Channel
	TransfersQeue    amqp.Queue
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

	q, _ := TransfersChannel.QueueDeclare(
		"transferencia",
		false,
		false,
		false,
		false,
		nil,
	)

	TransfersQeue = q
}
