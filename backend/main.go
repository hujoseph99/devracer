package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hujoseph99/typing/backend/db"
	"github.com/hujoseph99/typing/backend/graphql"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	db.InitDatabase()
	graphql.RegisterEndpoints(router)
	InitRouter(router)

	handler := cors.New(cors.Options{AllowedOrigins: []string{"http://localhost:3000"}, AllowCredentials: true}).Handler(router)

	port, check := os.LookupEnv("PORT")
	if !check {
		log.Fatal("No port env variable")
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
