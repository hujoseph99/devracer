package multiplayer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dchest/uniuri"
	"github.com/hujoseph99/typing/backend/common/utils"
	"github.com/hujoseph99/typing/backend/db"
)

// TODO: deal with event where someone joins after the game is finished... Might have to implement
// not allowing the user to race if the game is already in progress

type gameProgressData struct {
	client   *Client
	progress string
}

func newGameProgressData(client *Client, progress string) *gameProgressData {
	return &gameProgressData{
		client:   client,
		progress: progress,
	}
}

type Lobby struct {
	id            string
	snippet       *db.Snippet
	gameProgress  []*gameContent
	queuedPlayers []*gameContent
	placements    []string
	startTime     time.Time
	leader        *Client
	inProgress    bool
	hasReloaded   bool

	clients    map[*Client]bool
	start      chan bool
	register   chan *Client
	unregister chan *Client
	finisher   chan *Client
	progress   chan *gameProgressData
	startGame  chan *Client
	nextGame   chan *Client
	leaveGame  chan *Client
	broadcast  chan []byte
}

func generateLobbyId() string {
	return uniuri.NewLen(uniuri.StdLen)
}

func NewLobby(leader *Client) (*Lobby, error) {
	snippet, err := db.GetRandomSnippet(context.Background())
	if err != nil {
		return nil, err
	}
	res := &Lobby{
		id:            generateLobbyId(),
		snippet:       snippet,
		gameProgress:  make([]*gameContent, 0),
		queuedPlayers: make([]*gameContent, 0),
		placements:    make([]string, 0),
		startTime:     time.Now(),
		leader:        leader,
		inProgress:    false,
		hasReloaded:   true,

		clients:    make(map[*Client]bool),
		start:      make(chan bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		finisher:   make(chan *Client),
		progress:   make(chan *gameProgressData),
		startGame:  make(chan *Client),
		nextGame:   make(chan *Client),
		leaveGame:  make(chan *Client),
		broadcast:  make(chan []byte),
	}
	go res.RunLobby()
	return res, nil
}

func (lobby *Lobby) RunLobby() {
	for {
		select {
		case client := <-lobby.register:
			lobby.registerClientInLobby(client)
		case message := <-lobby.broadcast:
			lobby.broadcastToClientsInLobby(message)
		case client := <-lobby.finisher:
			lobby.handleFinisher(client)
		case <-lobby.start:
			lobby.inProgress = true
			lobby.startTime = time.Now()
		case progress := <-lobby.progress:
			lobby.handleGameProgress(progress)
		case client := <-lobby.startGame:
			lobby.handleStartGame(client)
		case client := <-lobby.nextGame:
			lobby.handleNextGame(client)
		case client := <-lobby.leaveGame:
			lobby.handleLeaveGame(client)
		}
	}
}

func (lobby *Lobby) registerClientInLobby(client *Client) {
	if _, exists := lobby.clients[client]; exists {
		createAndSendError(client, "You have already joined the lobby")
		return
	}
	// first send the client to all the players already in the lobby
	lobbyPayload := newNewPlayerResult(client.id, client.name, lobby.inProgress)
	lobbyResponse := newRequestResponse(newPlayerResponse, lobbyPayload)
	lobbyEncoded, err := json.Marshal(lobbyResponse)
	if err != nil {
		createAndSendError(client, "Error when trying to join the lobby")
		return
	}

	// in the case the bottom part errors, the ghost player would only be client side, so shouldn't
	// affect starting a new game
	lobby.broadcastToClientsInLobby(lobbyEncoded)

	// add the client to the lobby and then send all the details to the client
	lobby.clients[client] = true
	if lobby.inProgress {
		lobby.queuedPlayers = append(lobby.queuedPlayers, newGameContent(client.id, client.name))
	} else {
		lobby.gameProgress = append(lobby.gameProgress, newGameContent(client.id, client.name))
	}

	clientPayload := newJoinGameResult(client.id, lobby.snippet, lobby.gameProgress, lobby.queuedPlayers, lobby.placements, lobby.inProgress)
	clientResponse := newRequestResponse(joinGameResponse, clientPayload)
	clientEncoded, err := json.Marshal(clientResponse)
	if err != nil {
		// remove the client from the lobby so no ghost player shows up
		delete(lobby.clients, client)
		if len(lobby.gameProgress) > 0 {
			lobby.gameProgress = lobby.gameProgress[:len(lobby.gameProgress)-1]
		}

		createAndSendError(client, "Error when trying to join the lobby")
		return
	}
	client.send <- clientEncoded
}

func (lobby *Lobby) clientInPlacements(client *Client) bool {
	for _, val := range lobby.placements {
		if val == client.id {
			return true
		}
	}
	return false
}

func (lobby *Lobby) handleFinisher(client *Client) {
	// just do nothing if the client already in placements (will probably be unecessary, but just in case)
	if lobby.clientInPlacements(client) {
		return
	}

	lobby.placements = append(lobby.placements, client.id)

	payload := newPlayerFinishedResult(lobby.placements)
	response := newRequestResponse(playerFinishedResponse, payload)
	encoded, err := json.Marshal(response)
	if err != nil {
		createAndSendError(client, "There was an error when entering your placements.")
		return
	}
	lobby.broadcastToClientsInLobby(encoded)

	// check if game is finished
	if len(lobby.placements) == len(lobby.gameProgress) {
		gameFinishedPayload := newGameFinishedResult(lobby.placements)
		gameFinishedResponse := newRequestResponse(gameFinishedResponse, gameFinishedPayload)
		gameFinishedEncoded, err := json.Marshal(gameFinishedResponse)
		if err != nil {
			createAndSendErrorToLobby(lobby, "There was an error when returning your placements.")
			return
		}
		lobby.inProgress = false // finished, so now should let user start a new game
		lobby.broadcastToClientsInLobby(gameFinishedEncoded)
	}
}

func (lobby *Lobby) handleGameProgress(data *gameProgressData) {
	client := data.client

	if !lobby.inProgress {
		createAndSendError(client, "The race has not yet started.")
		return
	}
	// if the client already finished or not a part of the race, don't let them modify anything
	if _, err := lobby.findClientGameProgress(client); lobby.clientInPlacements(client) || err != nil {
		return
	}

	differenceIndex := utils.FindFirstDifference(lobby.snippet.RaceContent, data.progress)
	// this also checks if the user is indeed in the lobby
	gameProgress, err := lobby.findClientGameProgress(client)
	if err != nil {
		createAndSendError(client, "An error has occurred on the server. Please try rejoining the game.")
		return
	}

	difference := float64(differenceIndex) / float64(len(lobby.snippet.RaceContent))
	percentCompleted := utils.RoundFloor(difference, 2)
	secondsElapsed := time.Since(lobby.startTime).Round(time.Millisecond).Seconds()
	wpm := float64(differenceIndex) / 5.7 / secondsElapsed * 60 // average letters in a word is 4.7, so 5.7 including spaces
	gameProgress.PercentCompleted = percentCompleted
	gameProgress.Wpm = utils.RoundFloor(wpm, 0)

	gameProgressPayload := newGameProgressResult(client.id, gameProgress.PercentCompleted, gameProgress.Wpm)
	gameProgressResponse := newRequestResponse(gameProgressResponse, gameProgressPayload)
	gameProgressEncoded, err := json.Marshal(gameProgressResponse)
	if err != nil {
		createAndSendError(client, "An error occurred on the server. Please try rejoining the game.")
		return
	}
	lobby.broadcastToClientsInLobby(gameProgressEncoded)

	if percentCompleted == 1 && len(data.progress) == len(lobby.snippet.RaceContent) { // finished race
		lobby.handleFinisher(data.client)
	}
}

func (lobby *Lobby) handleStartGame(client *Client) {
	if lobby.leader != client {
		createAndSendError(client, "You are not the leader of the lobby.")
		return
	}
	if lobby.inProgress {
		createAndSendError(client, "The game is already in progress.")
		return
	}
	if !lobby.hasReloaded {
		createAndSendError(client, "Please wait for the leader to get a new game before starting.")
		return
	}

	lobby.hasReloaded = false

	for i := countdownStart; i >= 0; i-- {
		countdownCopy := i
		time.AfterFunc(time.Second*time.Duration(countdownStart-i), func() {
			if countdownCopy == 0 {
				lobby.start <- true // start the game
			}
			payload := newGameStartResult(countdownCopy)
			response := newRequestResponse(gameStartResponse, payload)
			encoded, err := json.Marshal(response)
			if err != nil {
				createAndSendError(client, "An error occurred on the server. There is an error sending the countdown")
				return
			}
			lobby.broadcast <- encoded
		})
	}
}

func (lobby *Lobby) handleNextGame(client *Client) {
	if lobby.leader != client {
		createAndSendError(client, "You are not the leader of the lobby.")
		return
	}
	if lobby.inProgress {
		createAndSendError(client, "The game is still in progress.")
		return
	}

	for _, val := range lobby.gameProgress {
		val.PercentCompleted = 0
		val.Wpm = 0
	}

	lobby.hasReloaded = true
	// add all the queued players over into the actual game
	lobby.gameProgress = append(lobby.gameProgress, lobby.queuedPlayers...)
	lobby.queuedPlayers = make([]*gameContent, 0)
	lobby.placements = make([]string, 0)
	newSnippet, err := db.GetRandomSnippet(context.Background())
	if err != nil {
		createAndSendError(client, "There was an error when trying to create a new game. Please try again.")
		return
	}
	lobby.snippet = newSnippet

	payload := newNextGameResult(lobby.snippet, lobby.gameProgress, lobby.queuedPlayers, lobby.placements)
	response := newRequestResponse(nextGameResponse, payload)
	encoded, err := json.Marshal(response)
	if err != nil {
		createAndSendError(client, "There was an error when trying to create a new game. Please try again.")
		return
	}
	lobby.broadcastToClientsInLobby(encoded)
}

func (lobby *Lobby) handleLeaveGame(client *Client) {
	if lobby.leader == client {
		payload := newLobbyClosedResult("The has been closed because the host left.")
		response := newRequestResponse(lobbyClosedResponse, payload)
		encoded, err := json.Marshal(response)
		if err != nil {
			log.Println("There was an error when closing the lobby")
		}
		lobby.broadcastToClientsInLobby(encoded)

		// remove the lobby from every client
		for clients := range lobby.clients {
			clients.lobby = nil
		}
		// delete the lobby
		client.server.delete <- lobby
	} else {
		// remove it from the map of clients
		delete(lobby.clients, client)
		// remove it from the slice of game progresses
		for progress := range lobby.gameProgress {
			if lobby.gameProgress[progress].PlayerId == client.id {
				lobby.gameProgress = append(lobby.gameProgress[:progress], lobby.gameProgress[progress+1:]...)
				break
			}
		}
		// remove it from the placements
		for placement := range lobby.placements {
			if lobby.placements[placement] == client.id {
				lobby.placements = append(lobby.placements[:placement], lobby.placements[placement+1:]...)
				break
			}
		}
		// remove it from the slice of queued players
		for progress := range lobby.queuedPlayers {
			if lobby.queuedPlayers[progress].PlayerId == client.id {
				lobby.queuedPlayers = append(lobby.queuedPlayers[:progress], lobby.queuedPlayers[progress+1:]...)
				break
			}
		}

		client.lobby = nil

		payload := newLeaveGameResult(client.id, lobby.placements)
		response := newRequestResponse(leaveGameResponse, payload)
		encoded, err := json.Marshal(response)
		if err != nil {
			log.Println("There was an error when encoding the leave game response")
			return
		}
		lobby.broadcastToClientsInLobby(encoded)
	}
}

func (lobby *Lobby) broadcastToClientsInLobby(message []byte) {
	for client := range lobby.clients {
		client.send <- message
	}
}

func (lobby *Lobby) findClientGameProgress(client *Client) (*gameContent, error) {
	for _, val := range lobby.gameProgress {
		if val.PlayerId == client.id {
			return val, nil
		}
	}

	return nil, fmt.Errorf("could not find client in the game")
}
