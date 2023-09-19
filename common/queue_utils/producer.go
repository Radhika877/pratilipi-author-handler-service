package queueutils

import (
	"context"
	"log"

	"author-handler-service/lib"
	queueModels "author-handler-service/models/queue"

	"github.com/streadway/amqp"
)

func Producer(config *lib.Config, payload queueModels.QueueStruct, ctx context.Context) {
	if payload.QueueName == "" || payload.Message.AuthorId == "" {
		log.Printf("Empty payload received with request ID %v", ctx.Value("REQUEST_ID"))
		return
	}
	log.Printf("Recived payload in producer %v with request ID %v", payload, ctx.Value("REQUEST_ID"))
	conn, err := amqp.Dial(config.AMQPEndpoint)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v with request ID %v", err, ctx.Value("REQUEST_ID"))
	}
	defer ch.Close()

	err = ch.Publish(
		"",
		payload.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload.Message.AuthorId),
		},
	)
	if err != nil {
		log.Printf("Failed to publish a message: %v with request ID %v", err, ctx.Value("REQUEST_ID"))
	}
	log.Printf("Message sent to consumer: %s with request ID %v", payload.Message.AuthorId, ctx.Value("REQUEST_ID"))
}
