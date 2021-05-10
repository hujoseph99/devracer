package multiplayer

import (
	"log"
	"net/http"

	"github.com/hujoseph99/typing/backend/common/api"
)

// handleCustomGame will handle upgrading the request to use websocket protocol
func HandleCustomGame(server *MultiplayerServer, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true // will have to check the origin in the future. For now, just enable all
	}

	name, ok := r.URL.Query()["name"]
	if !ok || len(name[0]) <= 0 {
		api.DefaultError(w, r, http.StatusBadRequest, "No name was provided")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, server, name[0])

	go client.writePump()
	go client.readPump()

	// server.register <- client
}
