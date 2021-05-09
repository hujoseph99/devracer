package multiplayer

type MultiplayerServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewMultiplayerServer() *MultiplayerServer {
	return &MultiplayerServer{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
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
