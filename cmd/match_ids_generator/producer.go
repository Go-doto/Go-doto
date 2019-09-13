package main

import (
	"encoding/json"
	"fmt"
	"github.com/Go-doto/Go-doto/internal"
	"github.com/spf13/viper"
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
	viper.AddConfigPath(os.Getenv("CONFIG_FILE_PATH"))
	viper.SetConfigName(os.Getenv("CONFIG_FILE_NAME"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
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

	var startNum int64 = 400000000
	var amount int = viper.GetInt("countOfMatchesToParse")
	task := internal.NewTask(startNum, amount)

	jsonTask, err := json.Marshal(&task)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonTask,
		})
	log.Printf(" [x] Sent %d", jsonTask)
	failOnError(err, "Failed to publish a message")
}
