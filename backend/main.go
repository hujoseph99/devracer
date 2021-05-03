package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hujoseph99/typing/backend/db"
	"github.com/hujoseph99/typing/backend/graphql"
)

func main() {
	router := mux.NewRouter()

	db.InitDatabase()
	graphql.RegisterEndpoints(router)
	InitRouter(router)
	http.ListenAndServe(":8080", router)
}
