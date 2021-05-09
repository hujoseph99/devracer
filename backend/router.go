package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hujoseph99/typing/backend/auth"
	"github.com/hujoseph99/typing/backend/multiplayer"
)

func InitRouter(router *mux.Router, server *multiplayer.MultiplayerServer) {
	router.HandleFunc("/auth/login", auth.HandleLogin)
	router.HandleFunc("/auth/register", auth.HandleRegister)
	router.HandleFunc("/auth/refresh", auth.HandleRefresh)
	router.HandleFunc("/auth/logout", auth.HandleLogout)
	router.HandleFunc("/custom", func(w http.ResponseWriter, r *http.Request) {
		multiplayer.HandleCustomGame(server, w, r)
	})
}
