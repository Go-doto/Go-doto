package main

import (
	"fmt"
	dotaApi "github.com/Go-doto/Go-doto/pkg/dota-api"
	"log"
	"os"
)

func main() {
	client, _ := dotaApi.NewClientWithToken("token")

	resp, err := dotaApi.GetMatchHistoryBySequenceNum(client, dotaApi.MatchSequenceNo(4151934576), 100)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	fmt.Printf("%+v\n", resp)
}
