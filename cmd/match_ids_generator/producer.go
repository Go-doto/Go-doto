package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Go-doto/Go-doto/internal"
	"github.com/gomodule/redigo/redis"
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
	forceStart := flag.Bool("forceStart", false, "Use forceStart to generate task even if worker queue empty")
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	workerQueue, err := ch.QueueDeclare(
		os.Getenv("WORKER_QUEUE"),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	parsingQueue, err := ch.QueueDeclare(
		os.Getenv("GAMES_PARSE_QUEUE_NAME"),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		workerQueue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")

	if *forceStart == true {
		sendTask(parsingQueue, *ch)
	}

	forever := make(chan bool)
	go func() {
		for range messages {
			sendTask(parsingQueue, *ch)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func sendTask(q amqp.Queue, ch amqp.Channel) {
	redisConnection, err := redis.Dial("tcp", os.Getenv("REDIS_URL"))
	if err != nil {
		failOnError(err, "Redis connection failed")
	}

	defer redisConnection.Close()

	key := viper.GetString("lastParsedMatchIdKey")
	startNum, err := redis.Int64(redisConnection.Do("GET", key))
	if err == redis.ErrNil {
		startNum = 1
	} else if err != nil {
		failOnError(err, "Failed to get last match id")
	}

	task := internal.ParseMatchTask{
		StartNum:     startNum,
		Amount:       viper.GetInt("countOfMatchesToParse"),
		QueriesLimit: viper.GetInt("queriesLimit"),
	}

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
