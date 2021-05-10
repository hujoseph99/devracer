package multiplayer

import "fmt"

type Lobby struct {
	id         string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
}

func NewLobby(id string) *Lobby {
	return &Lobby{
		id:         id,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
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
	lobby.notifyClientJoined(client)
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

const welcomeMessage = "%s joined the lobby"

func (lobby *Lobby) notifyClientJoined(client *Client) {
	message := &Message{
		Action:  SendMessageAction,
		Target:  lobby.id,
		Message: fmt.Sprintf(welcomeMessage, client.GetName()),
	}
	lobby.broadcastToClientsInLobby(message.encode())
}
