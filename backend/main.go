package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hujoseph99/typing/backend/db"
	"github.com/hujoseph99/typing/backend/graphql"
	"github.com/hujoseph99/typing/backend/multiplayer"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	db.InitDatabase()
	graphql.RegisterEndpoints(router)

	multiplayerServer := multiplayer.NewMultiplayerServer()
	go multiplayerServer.Run()

	InitRouter(router, multiplayerServer)

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}
