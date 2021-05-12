package multiplayer

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	createGameAction   = "createGame"
	joinGameAction     = "joinGame"
	gameProgressAction = "gameProgress"
	startGameAction    = "startGame"
	nextGameAction     = "nextGame"
	leaveGameAction    = "leaveGame"
)

type Message struct {
	Action  string `json:"action"`
	Payload string `json:"payload,omitempty"`
	LobbyId string `json:"lobbyId,omitempty"`
	client  *Client
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
