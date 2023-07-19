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
	"admin": "123456",
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
	log.Info("Handling error")
	w.WriteHeader(status)
}

func handleUnauthorized(w http.ResponseWriter) {
	log.Info("Unauthorized")
	handleError(w, http.StatusUnauthorized)
}

func handleBadRequest(w http.ResponseWriter) {
	log.Info("Bad request")
	handleError(w, http.StatusBadRequest)
}

func handleInternalError(w http.ResponseWriter) {
	log.Info("Internal server error")
	handleError(w, http.StatusInternalServerError)
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your middleware logic goes here. For example:
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

		// Call the next handler.
		next.ServeHTTP(w, r)
	})
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
		log.WithError(err).Error("Error generating the JWT token")
		handleInternalError(w)
		return
	}

	log.Info("Login successful")

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Refresh(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("token")

	if err != nil {
		log.WithError(err).Error("Error getting the cookie")
		if err == http.ErrNoCookie {
			log.WithError(err).Error("No cookie found")
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
		log.WithError(err).Error("Error parsing the JWT token")
		if err == jwt.ErrSignatureInvalid {
			log.WithError(err).Error("Invalid signature")
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

	log.Info("Received logout request")

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})

	log.Info("Logout successful")
}
