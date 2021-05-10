package multiplayer

import "fmt"

type MultiplayerServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	lobbies    map[*Lobby]bool
}

func NewMultiplayerServer() *MultiplayerServer {
	return &MultiplayerServer{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		lobbies:    make(map[*Lobby]bool),
	}
}

// Run our websocket server, accepting various requests
func (server *MultiplayerServer) Run() {
	for {
		select {
		case client := <-server.register:
			server.registerClient(client)
		case client := <-server.unregister:
			server.unregisterClient(client)
		case message := <-server.broadcast:
			server.broadcastToClients(message)
		}
	}
}

func (server *MultiplayerServer) registerClient(client *Client) {
	server.clients[client] = true
}

func (server *MultiplayerServer) unregisterClient(client *Client) {
	// apparently do not have to check if the key exists
	delete(server.clients, client)
}

func (server *MultiplayerServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

func (server *MultiplayerServer) findLobbyByID(id string) (*Lobby, error) {
	for lobby := range server.lobbies {
		if lobby.id == id {
			return lobby, nil
		}
	}
	return nil, fmt.Errorf("could not find lobby with id %s", id)
}

// have to make sure that the id is unique, otherwise will cause problems
func (server *MultiplayerServer) createLobby(id string) (*Lobby, error) {
	_, err := server.findLobbyByID(id)
	if err == nil {
		return nil, fmt.Errorf("a lobby with the given id was already created")
	}
	lobby := NewLobby(id)
	go lobby.RunLobby()
	server.lobbies[lobby] = true

	return lobby, nil
}
