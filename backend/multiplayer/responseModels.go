package multiplayer

import (
	"encoding/json"
	"log"

	"github.com/hujoseph99/typing/backend/db"
)

const (
	errorResponse          = "errorResponse"
	createGameReponse      = "createGameResponse"
	joinGameResponse       = "joinGameResponse"
	newPlayerResponse      = "newPlayerResponse"
	gameProgressResponse   = "gameProgressResponse"
	playerFinishedResponse = "playerFinishedResponse"
	gameFinishedResponse   = "gameFinishedResponse"
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

func createAndSendErrorToLobby(lobby *Lobby, message string) {
	log.Println(message)
	payload := newErrorResult(message)
	response := newRequestResponse(errorResponse, payload)
	encoded, err := json.Marshal(response)
	if err != nil {
		log.Println("error when handling error: ", err)
		return // in this case, silently do nothing and log the error LMAO
	}
	lobby.broadcastToClientsInLobby(encoded)
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
	PercentCompleted float64 `json:"percentCompleted"`
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
	Placements   []string               `json:"placements"`
}

func newJoinGameResult(playerId string, snippet *db.Snippet, gameProgress []*gameProgressContent,
	placements []string) *joinGameResult {
	return &joinGameResult{
		PlayerId:     playerId,
		Snippet:      snippet,
		GameProgress: gameProgress,
		Placements:   placements,
	}
}

type newPlayerResult struct {
	PlayerId         string  `json:"playerId"`
	DisplayName      string  `json:"displayName"`
	PercentCompleted float64 `json:"percentCompleted"`
}

func newNewPlayerResult(playerId string, displayName string) *newPlayerResult {
	return &newPlayerResult{
		PlayerId:         playerId,
		DisplayName:      displayName,
		PercentCompleted: 0,
	}
}

type gameProgressResult struct {
	PlayerId         string  `json:"playerId"`
	DisplayName      string  `json:"displayName"`
	PercentCompleted float64 `json:"percentCompleted"`
}

func newGameProgressResult(playerId string, displayName string, percentCompleted float64) *gameProgressResult {
	return &gameProgressResult{
		PlayerId:         playerId,
		DisplayName:      displayName,
		PercentCompleted: percentCompleted,
	}
}

type playerFinishedResult struct {
	Placements []string `json:"placements"`
}

func newPlayerFinishedResult(placements []string) *playerFinishedResult {
	return &playerFinishedResult{
		Placements: placements,
	}
}

type gameFinishedResult struct {
	Placements []string `json:"placements"`
}

func newGameFinishedResult(placements []string) *gameFinishedResult {
	return &gameFinishedResult{
		Placements: placements,
	}
}
