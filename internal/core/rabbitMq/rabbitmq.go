package rabbitmq

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

type RabbitMq struct {
	mqConnection     *amqp.Connection
	transfersChannel *amqp.Channel
	transfersQueue   amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func NewRabbitMq() RabbitMq {
	godotenv.Load()

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")

	channel, errChannel := conn.Channel()
	failOnError(errChannel, "Failed to connect to RabbitMQ")

	queue, errQueue := channel.QueueDeclare(
		"transferencia",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(errQueue, "Failed to connect to RabbitMQ")

	return RabbitMq{
		mqConnection:     conn,
		transfersChannel: channel,
		transfersQueue:   queue,
	}
}

func (r RabbitMq) PublishTransaction(transaction entity.Transaction, newBalance int) {
	data := struct {
		Transaction entity.Transaction
		NewBalance  int
	}{
		Transaction: transaction,
		NewBalance:  newBalance,
	}
	body, _ := json.Marshal(data)
	err := r.transfersChannel.Publish(
		"",
		r.transfersQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	failOnError(err, "Failed to publish a message")
}

func (r RabbitMq) ConsumeMessages(worker func(body []byte, context context.Context) error) {
	r.transfersChannel.Qos(1, 0, false)
	msgs, err := r.transfersChannel.Consume(
		r.transfersQueue.Name,
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
