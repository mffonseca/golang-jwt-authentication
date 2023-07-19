package main

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	claims := &Claims{}

	tknStr, _ := r.Cookie("token")

	_, err := jwt.ParseWithClaims(tknStr.Value, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.WithError(err).Error("Failed to parse token claims")
		handleInternalError(w)
		return
	}

	w.Write([]byte("Welcome " + claims.Username + "!"))

	log.Info("Welcome successful")
}
