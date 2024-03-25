package rabbitMq

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"os"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/entity"
	"github.com/joho/godotenv"
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
	godotenv.Load()

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
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

func PublishTransaction(transaction entity.Transaction, newBalance int) {
	data := struct {
		Transaction entity.Transaction
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

func ConsumeMessages(worker func(body []byte, context context.Context) error) {
	TransfersChannel.Qos(1, 0, false)
	msgs, err := TransfersChannel.Consume(
		TransfersQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")
	slog.Info("Wating messages...")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			err := worker(d.Body, context.Background())
			if err == nil {
				d.Ack(true)
			}
		}
	}()

	<-forever
}
