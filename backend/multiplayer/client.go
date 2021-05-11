package multiplayer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dchest/uniuri"
	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// Client represents the websocket client at the server
type Client struct {
	id     string
	name   string
	conn   *websocket.Conn
	server *MultiplayerServer
	send   chan []byte
	lobby  *Lobby
}

func generateClientId() string {
	return uniuri.NewLen(uniuri.StdLen)
}

func newClient(conn *websocket.Conn, server *MultiplayerServer, name string) *Client {
	return &Client{
		id:     generateClientId(),
		name:   name,
		conn:   conn,
		server: server,
		send:   make(chan []byte, 256),
		lobby:  nil,
	}
}

func (client *Client) GetName() string {
	return client.name
}

func (client *Client) disconnect() {
	if client.lobby != nil {
		client.lobby.unregister <- client
	}
	client.conn.Close()
}

func (client *Client) handleNewMessage(jsonMessage []byte) error {
	message, err := decode(jsonMessage, client)
	if err != nil {
		log.Printf("error on unmarshal JSON message %s", err)
		return fmt.Errorf("error on unmarshal JSON message %s", err)
	}

	switch message.Action {
	case createGameAction:
		client.handleCreateGameAction()
	case joinGameAction:
		client.handleJoinGameAction(message)
	}

	return nil

	// switch message.Action {
	// case SendMessageAction:
	// 	if client.lobby != nil {
	// 		client.lobby.broadcast <- &message
	// 	}
	// case JoinRoomAction:
	// 	client.handleJoinRoomMessage(message)
	// case LeaveRoomAction:
	// 	client.handleLeaveRoomMessage(message)
	// }
}

// can ignore the payload in this case and lobby id
func (client *Client) handleCreateGameAction() {
	var err error
	client.lobby, err = NewLobby()
	if err != nil {
		log.Println(err)
		return
		// TODO: handle error
	}

	client.server.create <- client.lobby // add the lobby to the map
	client.lobby.register <- client

	payload := newCreateGameResult(client.id, client.lobby.id, client.lobby.snippet)
	response := newRequestResponse(createGameReponse, payload)
	encoded, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
		// TODO: handle error
	}

	client.send <- encoded
}

// can ignore the payload in this case
func (client *Client) handleJoinGameAction(message *Message) {
	if len(message.LobbyId) <= 0 {
		log.Println("invalid lobby id")
		return
		// TODO: handle error
	}
	lobby, err := client.server.findLobbyByID(message.LobbyId)
	if err != nil {
		log.Println("invalid lobby id")
		return
		// TODO: handle error
	}

	lobby.register <- client
}

// func (client *Client) handleJoinRoomMessage(message Message) {
// 	lobbyID := message.Target

// 	lobby, err := client.server.findLobbyByID(lobbyID)
// 	if err != nil {
// 		lobby, _ = client.server.createLobby(lobbyID)
// 	}

// 	client.lobby = lobby
// 	lobby.register <- client
// }

// func (client *Client) handleLeaveRoomMessage(message Message) {
// 	room, err := client.server.findLobbyByID(message.Target)
// 	client.lobby = nil
// 	if err != nil {
// 		room.unregister <- client
// 	}
// }

func (client *Client) readPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error {
		client.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}
		client.handleNewMessage(jsonMessage)
		// if client.lobby != nil {
		// 	fmt.Println("-")
		// 	fmt.Println(client.lobby.clients)
		// 	fmt.Println(client.lobby.gameProgress)
		// }
		fmt.Println(client.server.lobbies)
	}
}

func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		client.disconnect()
	}()

	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// the server closed the channel
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
