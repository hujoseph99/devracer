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

	countdownStart = 5 // time to start the countdown from
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

func (client *Client) handleNewMessage(jsonMessage []byte) {
	message, err := decode(jsonMessage, client)
	if err != nil {
		createAndSendError(client, "Invalid message was sent")
		return
	}

	switch message.Action {
	case createGameAction:
		client.handleCreateGameAction()
	case joinGameAction:
		client.handleJoinGameAction(message)
	case gameProgressAction:
		client.handleGameProgressAction(message)
	case startGameAction:
		client.handleGameStartAction()
	case nextGameAction:
		client.handleNextGameAction()
	default:
		createAndSendError(client, "Invalid message was sent")
	}
}

// can ignore the payload in this case and lobby id
func (client *Client) handleCreateGameAction() {
	var err error
	client.lobby, err = NewLobby(client)
	if err != nil {
		createAndSendError(client, "An error occurred when creating a lobby")
		return
	}

	client.server.create <- client.lobby // add the lobby to the map
	client.lobby.register <- client

	payload := newCreateGameResult(client.id, client.lobby.id, client.lobby.snippet)
	response := newRequestResponse(createGameReponse, payload)
	encoded, err := json.Marshal(response)
	if err != nil {
		// do not have to delete the created lobby here. Let pseudo-garbage collector deal with it.
		createAndSendError(client, "An error occurred when joining the created lobby, please try again.")
		return
	}

	client.send <- encoded
}

// can ignore the payload in this case
func (client *Client) handleJoinGameAction(message *Message) {
	if len(message.LobbyId) <= 0 {
		createAndSendError(client, "An invalid lobby id was provided")
		return
	}
	lobby, err := client.server.findLobbyByID(message.LobbyId)
	if err != nil {
		createAndSendError(client, "An invalid lobby id was provided")
		return
	}

	client.lobby = lobby
	lobby.register <- client
}

// we expect the payload to be the current string that the user has typed in. Can ignore the lobbyId
// now because they SHOULD be in a lobby if they send this
func (client *Client) handleGameProgressAction(message *Message) {
	lobby := client.lobby
	if lobby == nil {
		createAndSendError(client, "You have not joined a lobby.")
		return
	}
	progress := message.Payload
	progressData := newGameProgressData(client, progress)
	client.lobby.progress <- progressData
}

// can more or less ignore the message here
func (client *Client) handleGameStartAction() {
	lobby := client.lobby
	if lobby == nil {
		createAndSendError(client, "You have not joined a lobby.")
		return
	}
	lobby.startGame <- client
}

func (client *Client) handleNextGameAction() {
	lobby := client.lobby
	if lobby == nil {
		createAndSendError(client, "You have not joined a lobby.")
		return
	}
	client.lobby.nextGame <- client
}

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
		fmt.Println("--------------")
		if client.lobby != nil {
			fmt.Println(client.lobby.clients)
			fmt.Println(client.lobby.gameProgress)
		}
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
