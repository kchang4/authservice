package main

import (
	"auth/internal/signup"

	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()
	authRouter := mainRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", signup.SignupHandler).Methods("GET")
}
