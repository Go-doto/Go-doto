package main

import (
	"errors"
	"fmt"
	dotaApi "github.com/Go-doto/Go-doto/pkg/dota-api"
	"log"
	"os"
)

func main() {
	client, _ := dotaApi.NewClientWithToken("token")

	resp, err := dotaApi.GetMatchHistoryBySequenceNum(client, dotaApi.MatchSequenceNo(4151934576), 100)
	if errors.As(err, &dotaApi.AccessForbiddenError{}) {
		log.Fatal("invalid token")
	}
	if errors.As(err, &dotaApi.ValidationError{}) {
		log.Fatalf("validation error: %s", err)
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	fmt.Printf("%+v\n", resp)
}
