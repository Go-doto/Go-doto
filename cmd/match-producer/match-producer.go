package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
	"time"
)

const topic = "New"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn, err := kafka.DialLeader(context.Background(), os.Getenv("KAFKA_NETWORK"), os.Getenv("KAFKA_HOST"), topic, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {

		err = conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		if err != nil {
			log.Fatal(err)
		}
		n := time.Now().Unix()
		// write timestamp to queue
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(strconv.FormatInt(n, 10))},
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Written to topic:", n)

		time.Sleep(time.Second)

	}

}
