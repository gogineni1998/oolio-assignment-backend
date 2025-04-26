package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gogineni1998/oolio-assignment-backend/configuration"
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

		expectedPassword, ok := models.Users[creds.Username]
		if !ok || expectedPassword != creds.Password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		expirationTime := time.Now().Add(15 * time.Minute)
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
		_, ok := models.Users[creds.Username]
		if ok {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("user already exists"))
			return
		}
		models.Users[creds.Username] = creds.Password
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("user registered successfully"))
	}
}
