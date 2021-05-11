package multiplayer

import (
	"context"
	"encoding/json"

	"github.com/dchest/uniuri"
	"github.com/hujoseph99/typing/backend/db"
)

type Lobby struct {
	id           string
	snippet      *db.Snippet
	gameProgress []*gameProgressContent
	clients      map[*Client]bool
	register     chan *Client
	unregister   chan *Client
	broadcast    chan []byte
}

func generateLobbyId() string {
	return uniuri.NewLen(uniuri.StdLen)
}

func NewLobby() (*Lobby, error) {
	snippet, err := db.GetRandomSnippet(context.Background())
	if err != nil {
		return nil, err
	}
	res := &Lobby{
		id:           generateLobbyId(),
		snippet:      snippet,
		gameProgress: make([]*gameProgressContent, 0),
		clients:      make(map[*Client]bool),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		broadcast:    make(chan []byte),
	}
	go res.RunLobby()
	return res, nil
}

func (lobby *Lobby) RunLobby() {
	for {
		select {
		case client := <-lobby.register:
			lobby.registerClientInLobby(client)
		// case client := <-lobby.unregister:
		// 	lobby.unregisterClientInLobby(client)
		case message := <-lobby.broadcast:
			lobby.broadcastToClientsInLobby(message)
		}
	}
}

func (lobby *Lobby) registerClientInLobby(client *Client) {
	// first send the client to all the players already in the lobby
	lobbyPayload := newNewPlayerResult(client.id, client.name)
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
	lobby.gameProgress = append(lobby.gameProgress, newGameProgressContent(client.id, client.name))

	clientPayload := newJoinGameResult(client.id, lobby.snippet, lobby.gameProgress)
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

// func (lobby *Lobby) unregisterClientInLobby(client *Client) {
// 	if _, ok := lobby.clients[client]; ok {
// 		delete(lobby.clients, client)
// 	}
// }

func (lobby *Lobby) broadcastToClientsInLobby(message []byte) {
	for client := range lobby.clients {
		client.send <- message
	}
}
