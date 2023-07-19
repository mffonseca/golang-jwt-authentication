package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})

	log.SetLevel(log.DebugLevel)

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error getting environment variables")
	}

	router := mux.NewRouter()

	router.HandleFunc("/signin", Signin).Methods("POST")
	router.HandleFunc("/welcome", Welcome).Methods("GET")
	router.HandleFunc("/refresh", Refresh).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
