package cmd

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"time"
)

const topic = "New"

// helloCmd represents the hello command
var matchProducerCmd = &cobra.Command{
	Use:   "match-producer",
	Short: "Run match history producer",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runProducer()
	},
}

func runProducer() {
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

func init() {
	RootCmd.AddCommand(matchProducerCmd)
}
