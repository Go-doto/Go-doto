package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"sync"
)

const inputTopic = "New"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//  Kafka subscriber (consumer)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST")},
		Topic:   inputTopic,
		// GroupID:   "consumer-group-id-3",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	var wg sync.WaitGroup
	for {
		ctx := context.Background()

		// Read messages from kafka
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println(err)
			break
		}

		wg.Add(1)
		go printer(m)

	}
	wg.Wait()

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

}

func printer(m kafka.Message) {
	fmt.Println(string(m.Value))
}
