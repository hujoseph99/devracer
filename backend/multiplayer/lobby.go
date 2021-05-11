package multiplayer

import (
	"context"

	"github.com/dchest/uniuri"
	"github.com/hujoseph99/typing/backend/db"
)

type Lobby struct {
	id         string
	snippet    *db.Snippet
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
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
		id:         generateLobbyId(),
		snippet:    snippet,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
	go res.RunLobby()
	return res, nil
}

func (lobby *Lobby) RunLobby() {
	for {
		select {
		case client := <-lobby.register:
			lobby.registerClientInLobby(client)
		case client := <-lobby.unregister:
			lobby.unregisterClientInLobby(client)
		case message := <-lobby.broadcast:
			lobby.broadcastToClientsInLobby(message.encode())
		}
	}
}

func (lobby *Lobby) registerClientInLobby(client *Client) {
	lobby.clients[client] = true
}

func (lobby *Lobby) unregisterClientInLobby(client *Client) {
	if _, ok := lobby.clients[client]; ok {
		delete(lobby.clients, client)
	}
}

func (lobby *Lobby) broadcastToClientsInLobby(message []byte) {
	for client := range lobby.clients {
		client.send <- message
	}
}
