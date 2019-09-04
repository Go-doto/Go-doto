package dota_parser

import (
	"fmt"
	DotaApi "github.com/Go-doto/Go-doto/pkg/dota-api"
	"log"
	"os"
	"strconv"
)

func HandleGame(id GameId) {
	//token in env
	client, _ := DotaApi.NewClientWithToken(os.Getenv("STEAM_API_KEY"))

	//TODO::discuss type conversions
	resp, err := DotaApi.GetMatchDetails(client, strconv.Itoa(int(id)))
	if err != nil {
		log.Fatal(err)
	}
	//decode resp in struct
	//use repo collection to save
	fmt.Println(resp)
}
