package dota_parser

import (
	"fmt"
	DotaApi "github.com/Go-doto/Go-doto/pkg/dota-api"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)

func HandleGame(id []byte) {
	matchId, err := DotaApi.CreateMatchId(id)
	client, _ := DotaApi.NewClientWithToken(os.Getenv("STEAM_API_KEY"))
	resp, err := DotaApi.GetMatchDetails(client, matchId)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := redis.Dial("tcp", os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err)
	}
	// Importantly, use defer to ensure the connection is always
	// properly closed before exiting the main() function.
	defer conn.Close()

	//decode resp in struct
	//use repo collection to save
	fmt.Println(resp)
}
