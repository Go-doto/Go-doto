package dota_parser

import (
	"encoding/json"
	"fmt"
	"github.com/Go-doto/Go-doto/internal"
	DotaApi "github.com/Go-doto/Go-doto/pkg/dota-api"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func init() {
	gotenv.Load()
	viper.AddConfigPath(os.Getenv("CONFIG_FILE_PATH"))
	viper.SetConfigName(os.Getenv("CONFIG_FILE_NAME"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func HandleGame(task []byte) {
	redisConnection, err := redis.Dial("tcp", os.Getenv("REDIS_URL"))
	defer redisConnection.Close()

	if err != nil {
		log.Fatal(err)
	}

	key := viper.GetString("lastParsedMatchIdKey")
	var parseMatchTask internal.ParseMatchTask
	unmarshalError := json.Unmarshal(task, &parseMatchTask)
	if unmarshalError != nil {
		log.Fatal(err)
	}

	client, httpClientError := DotaApi.NewClientWithToken(os.Getenv("STEAM_API_KEY"))
	if httpClientError != nil {
		log.Fatal(httpClientError)
	}

	startNum := DotaApi.MatchSequenceNo(parseMatchTask.StartNum)
	for i := 0; i <= parseMatchTask.QueriesLimit; i++ {
		resp, err := DotaApi.GetMatchHistoryBySequenceNum(client, startNum, parseMatchTask.Amount)
		if err != nil {
			log.Fatal(err)
		}
		redisSetError := redisConnection.Send("SET", key, resp.MatchesResult[len(resp.MatchesResult)-1].MatchSequenceNo)
		if redisSetError != nil {
			log.Fatal(redisSetError)
		}
		//Хуёво выглядит
		startNum = DotaApi.MatchSequenceNo(int64(startNum) + int64(parseMatchTask.Amount))
	}
}
