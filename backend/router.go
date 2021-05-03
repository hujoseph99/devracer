package main

import (
	"github.com/gorilla/mux"
	"github.com/hujoseph99/typing/backend/auth"
)

func InitRouter(router *mux.Router) {
	router.HandleFunc("/auth/login", auth.HandleLogin)
}