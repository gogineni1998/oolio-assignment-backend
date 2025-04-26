package handlers

import (
	"fmt"
	"net/http"

	"github.com/gogineni1998/oolio-assignment-backend/configuration"
)

var Product = func() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value(configuration.UsernameContextKey)
		if r.Method != "GET" {
			http.Error(w, "Invalid request type", http.StatusForbidden)
			return
		}
		fmt.Println("Product handler called", username.(string))
	}
}
