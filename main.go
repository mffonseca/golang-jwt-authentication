package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const serverPort = ":8080"

func main() {

	// Set the log format to JSON
	log.SetFormatter(&log.JSONFormatter{})

	// Set the log level to Debug
	log.SetLevel(log.DebugLevel)

	// Load environment variables from .env file
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Error getting environment variables: ", errEnv)
	}

	// Initialize router
	router := mux.NewRouter()

	// Setup routes
	router.HandleFunc("/signin", Signin).Methods("POST")
	router.Handle("/welcome", AuthenticationMiddleware(http.HandlerFunc(Welcome))).Methods("GET")
	router.HandleFunc("/refresh", Refresh).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("GET")

	// Start the server
	errListen := http.ListenAndServe(serverPort, router)

	// If there was an error starting the server, log it and exit
	if errListen != nil {
		log.Fatalf("There was an error starting the server: %v", errListen)
	}

	log.Info("Server started on port " + serverPort)

}
