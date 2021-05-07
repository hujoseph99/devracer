package main

import (
	"github.com/gorilla/mux"
	"github.com/hujoseph99/typing/backend/auth"
)

func InitRouter(router *mux.Router) {
	router.HandleFunc("/auth/login", auth.HandleLogin)
	router.HandleFunc("/auth/register", auth.HandleRegister)
	router.HandleFunc("/auth/refresh", auth.HandleRefresh)
	router.HandleFunc("/auth/logout", auth.HandleLogout)
}
