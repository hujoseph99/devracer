package multiplayer

import (
	"encoding/json"
	"log"

	"github.com/hujoseph99/typing/backend/db"
)

const (
	errorResponse     = "errorResponse"
	createGameReponse = "createGameResponse"
	joinGameResponse  = "joinGameResponse"
	newPlayerResponse = "newPlayerResponse"
)

type requestResponse struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

func newRequestResponse(action string, payload interface{}) *requestResponse {
	return &requestResponse{
		Action:  action,
		Payload: payload,
	}
}

type errorResult struct {
	Message string `json:"message"`
}

func newErrorResult(message string) *errorResult {
	return &errorResult{
		Message: message,
	}
}

func createAndSendError(client *Client, message string) {
	log.Println(message)
	payload := newErrorResult(message)
	response := newRequestResponse(errorResponse, payload)
	encoded, err := json.Marshal(response)
	if err != nil {
		log.Println("error when handling error: ", err)
		return // in this case, silently do nothing and log the error LMAO
	}
	client.send <- encoded
}

type createGameResult struct {
	PlayerId string      `json:"playerId"`
	LobbyId  string      `json:"lobbyId"`
	Snippet  *db.Snippet `json:"snippet"`
}

func newCreateGameResult(playerId string, lobbyId string, snippet *db.Snippet) *createGameResult {
	return &createGameResult{
		PlayerId: playerId,
		LobbyId:  lobbyId,
		Snippet:  snippet,
	}
}

type gameProgressContent struct {
	PlayerId         string  `json:"playerId"`
	DisplayName      string  `json:"displayName"`
	PercentCompleted float32 `json:"percentCompleted"`
}

func newGameProgressContent(playerId string, displayName string) *gameProgressContent {
	return &gameProgressContent{
		PlayerId:         playerId,
		DisplayName:      displayName,
		PercentCompleted: 0,
	}
}

type joinGameResult struct {
	PlayerId     string                 `json:"playerId"`
	Snippet      *db.Snippet            `json:"snippet"`
	GameProgress []*gameProgressContent `json:"gameProgress"`
}

func newJoinGameResult(playerId string, snippet *db.Snippet, gameProgress []*gameProgressContent) *joinGameResult {
	return &joinGameResult{
		PlayerId:     playerId,
		Snippet:      snippet,
		GameProgress: gameProgress,
	}
}

type newPlayerResult struct {
	PlayerId         string  `json:"playerId"`
	DisplayName      string  `json:"displayName"`
	PercentCompleted float32 `json:"percentCompleted"`
}

func newNewPlayerResult(playerId string, displayName string) *newPlayerResult {
	return &newPlayerResult{
		PlayerId:         playerId,
		DisplayName:      displayName,
		PercentCompleted: 0,
	}
}

type gameProgressResult struct {
	GameProgress []*gameProgressContent `json:"gameProgress"`
}
