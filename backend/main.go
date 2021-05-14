package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hujoseph99/typing/backend/db"
	"github.com/hujoseph99/typing/backend/graphql"
	"github.com/hujoseph99/typing/backend/multiplayer"
	"github.com/hujoseph99/typing/backend/secret"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	db.InitDatabase()
	graphql.RegisterEndpoints(router)

	multiplayerServer := multiplayer.NewMultiplayerServer()
	go multiplayerServer.RunServer()

	InitRouter(router, multiplayerServer)

	handler := cors.New(cors.Options{AllowedOrigins: []string{secret.FrontendHostname}, AllowCredentials: true}).Handler(router)

	port, check := os.LookupEnv("PORT")
	if !check {
		log.Fatal("No port env variable")
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
