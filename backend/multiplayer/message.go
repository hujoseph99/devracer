package multiplayer

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	createGameAction   = "createGame"
	joinGameAction     = "joinGame"
	startGameAction    = "startGame"
	gameProgressAction = "gameProgress"
	nextGameAction     = "nextGame"
)

type Message struct {
	Action  string `json:"action"`
	Payload string `json:"payload,omitempty"`
	LobbyId string `json:"lobbyId,omitempty"`
	client  *Client
}

func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return json
}

func decode(jsonMessage []byte, client *Client) (*Message, error) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("error on unmarshal JSON message %s", err)
		return nil, fmt.Errorf("error on unmarshal JSON message %s", err)
	}

	message.client = client
	return &message, nil
}
