package main

import (
	DotaParser "github.com/Go-doto/Go-doto/pkg/dota-parser"
	"github.com/streadway/amqp"
	"github.com/subosito/gotenv"
	"log"
	"os"
	"time"
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

	parsingQueue, err := ch.QueueDeclare(
		os.Getenv("GAMES_PARSE_QUEUE_NAME"),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	workerQueue, err := ch.QueueDeclare(
		os.Getenv("WORKER_QUEUE"),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")

	messages, err := ch.Consume(
		parsingQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)
			DotaParser.HandleGame(message.Body)
			err = ch.Publish(
				"",
				workerQueue.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        make([]byte, 0),
				})
			time.Sleep(time.Minute)
			log.Printf("Done")
			message.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
