package multiplayer

import (
	"testing"
)

func TestCreateAndDeleteLobby(t *testing.T) {
	server := NewMultiplayerServer()
	lobby := server.createLobby()
	foundLobby, err := server.findLobbyByID(lobby.id)
	if err != nil {
		t.Fatal("could not find lobby")
	}
	if foundLobby.id != lobby.id {
		t.Fatal("found incorrect lobby")
	}
	server.deleteLobby(lobby.id)
	_, err = server.findLobbyByID(lobby.id)
	if err == nil {
		t.Fatal("found lobby")
	}
}
