package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	gotenv.Load()
}

func main() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("GAMES_PARSE_QUEUE_NAME"),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	//code below should be generator :D
	gameId, err := json.Marshal(4000000124)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        gameId,
		})
	log.Printf(" [x] Sent %d", gameId)
	failOnError(err, "Failed to publish a message")
}
