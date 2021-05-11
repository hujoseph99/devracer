package multiplayer

import (
	"context"
	"encoding/json"
	"log"

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
	lobby.clients[client] = true
	lobby.gameProgress = append(lobby.gameProgress, newGameProgressContent(client.id, client.name))

	payload := newJoinGameResult(client.id, lobby.snippet, lobby.gameProgress)
	response := newRequestResponse(joinGameResponse, payload)
	encoded, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
		// TODO: handle error
	}

	lobby.broadcastToClientsInLobby(encoded)
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
