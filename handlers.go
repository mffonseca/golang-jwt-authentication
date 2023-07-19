package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func handleError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func handleUnauthorized(w http.ResponseWriter) {
	handleError(w, http.StatusUnauthorized)
}

func handleBadRequest(w http.ResponseWriter) {
	handleError(w, http.StatusBadRequest)
}

func handleInternalError(w http.ResponseWriter) {
	handleError(w, http.StatusInternalServerError)
}

func Signin(w http.ResponseWriter, r *http.Request) {

	log.Info("Received login request")

	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		log.WithError(err).Error("Erro ao decodificar o corpo JSON")
		handleBadRequest(w)
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		handleUnauthorized(w)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.WithError(err).Error("Erro ao gerar o token JWT")
		handleInternalError(w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			handleUnauthorized(w)
			return
		}
		handleBadRequest(w)
		return
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			handleUnauthorized(w)
			return
		}
		handleBadRequest(w)
		return
	}

	if !tkn.Valid {
		handleUnauthorized(w)
		return
	}

	w.Write([]byte("Welcome " + claims.Username + "!"))
}

func Refresh(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			handleUnauthorized(w)
			return
		}
		handleBadRequest(w)
		return
	}

	tknStr := c.Value
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			handleUnauthorized(w)
			return
		}
		handleBadRequest(w)
		return
	}

	if !tkn.Valid {
		handleUnauthorized(w)
		return
	}

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		handleBadRequest(w)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		handleInternalError(w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}
