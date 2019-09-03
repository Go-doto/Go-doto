package main

import (
	"fmt"
	dotaApi "github.com/Go-doto/Go-doto/pkg/dota-api"
	"log"
	"os"
)

func main() {
	client, _ := dotaApi.NewClientWithToken("tokenstring")

	resp, err := dotaApi.GetMatchDetails(client, "4949341670")
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	fmt.Println(resp)
}
