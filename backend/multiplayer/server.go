package multiplayer

import (
	"fmt"
)

// TODO: Have to periodically go through the lobbies and check to see if any are inactive.
// 	If any have been inactive for more than 5 minutes, then just close the lobby.
// TODO: Batch responses to frontend, race progress response
// TODO: Add WPM

type MultiplayerServer struct {
	create  chan *Lobby
	delete  chan *Lobby
	lobbies map[string]*Lobby
}

func NewMultiplayerServer() *MultiplayerServer {
	return &MultiplayerServer{
		create:  make(chan *Lobby),
		delete:  make(chan *Lobby),
		lobbies: make(map[string]*Lobby),
	}
}

func (server *MultiplayerServer) RunServer() {
	for {
		select {
		case lobby := <-server.create:
			server.addLobby(lobby)
		case lobby := <-server.delete:
			server.deleteLobby(lobby.id)
		}
	}
}

func (server *MultiplayerServer) findLobbyByID(id string) (*Lobby, error) {
	if lobby, ok := server.lobbies[id]; ok {
		return lobby, nil
	}
	return nil, fmt.Errorf("could not find lobby with id %s", id)
}

func (server *MultiplayerServer) addLobby(lobby *Lobby) {
	server.lobbies[lobby.id] = lobby
}

func (server *MultiplayerServer) deleteLobby(id string) {
	delete(server.lobbies, id)
}
