package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gogineni1998/oolio-assignment-backend/configuration"
	"github.com/gogineni1998/oolio-assignment-backend/database"
	"github.com/gogineni1998/oolio-assignment-backend/models"
	"github.com/golang-jwt/jwt/v5"
)

var Token = func() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request type", http.StatusForbidden)
			return
		}
		var creds models.Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		credentials, err := database.CheckUser(configuration.DBUsersCollection, creds.Username)
		if err != nil {
			http.Error(w, "Error checking user", http.StatusInternalServerError)
			return
		}
		if credentials == nil {
			http.Error(w, "User dosen't exists", http.StatusConflict)
			return
		}
		if credentials.Password != creds.Password {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(60 * time.Minute)
		claims := &models.Claims{
			Username: creds.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
				Issuer:    "oolio-idp",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(configuration.JwtKey)
		if err != nil {
			http.Error(w, "Could not create token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

var Register = func() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request type", http.StatusForbidden)
			return
		}
		var creds models.Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		credentials, err := database.CheckUser(configuration.DBUsersCollection, creds.Username)
		if err != nil {
			http.Error(w, "Error checking user", http.StatusInternalServerError)
			return
		}
		if credentials != nil {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
		_, err = database.InsertUser(configuration.DBUsersCollection, creds)
		if err != nil {
			http.Error(w, "Error inserting user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("user registered successfully"))
	}
}
