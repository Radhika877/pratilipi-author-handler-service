package queueutils

import (
	"log"

	"author-handler-service/lib"
	queueModels "author-handler-service/models/queue"

	"github.com/streadway/amqp"
)

func Producer(config *lib.Config, payload queueModels.QueueStruct) {
	conn, err := amqp.Dial(config.AMQPEndpoint)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
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
		log.Printf("Failed to publish a message: %v", err)
	}
	log.Printf("Message sent to consumer:%s", payload.Message.AuthorId)
}
