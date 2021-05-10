package multiplayer

import (
	"fmt"

	"github.com/dchest/uniuri"
)

type MultiplayerServer struct {
	lobbies map[string]*Lobby
}

func NewMultiplayerServer() *MultiplayerServer {
	return &MultiplayerServer{
		lobbies: make(map[string]*Lobby),
	}
}

func generateLobbyId() string {
	return uniuri.NewLen(uniuri.StdLen)
}

func (server *MultiplayerServer) findLobbyByID(id string) (*Lobby, error) {
	if lobby, ok := server.lobbies[id]; ok {
		return lobby, nil
	}
	return nil, fmt.Errorf("could not find lobby with id %s", id)
}

func (server *MultiplayerServer) createLobby() *Lobby {
	id := generateLobbyId()
	lobby := NewLobby(id)

	go lobby.RunLobby()

	server.lobbies[id] = lobby

	return lobby
}

func (server *MultiplayerServer) deleteLobby(id string) {
	delete(server.lobbies, id)
}
