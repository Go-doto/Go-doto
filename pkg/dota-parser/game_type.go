package dota_parser

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type GameId int

func DecodeGameId(rawMessage amqp.Delivery) (GameId, error) {
	var gameId GameId
	err := json.Unmarshal(rawMessage.Body, &gameId)

	return gameId, err
}
