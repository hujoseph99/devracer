package multiplayer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// Client represents the websocket client at the server
type Client struct {
	Name   string `json:"name"`
	conn   *websocket.Conn
	server *MultiplayerServer
	send   chan []byte
	lobby  *Lobby
}

func newClient(conn *websocket.Conn, server *MultiplayerServer, name string) *Client {
	return &Client{
		Name:   name,
		conn:   conn,
		server: server,
		send:   make(chan []byte, 256),
		lobby:  nil,
	}
}

func (client *Client) GetName() string {
	return client.Name
}

func (client *Client) disconnect() {
	if client.lobby != nil {
		client.lobby.unregister <- client
	}
	client.conn.Close()
}

func (client *Client) handleNewMessage(jsonMessage []byte) error {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("error on unmarshal JSON message %s", err)
		return fmt.Errorf("error on unmarshal JSON message %s", err)
	}

	message.Sender = client

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

	return nil
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
		fmt.Println(string(jsonMessage))
		client.handleNewMessage(jsonMessage)
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

			// attach queued chat messages to the current websocket message.
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

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
