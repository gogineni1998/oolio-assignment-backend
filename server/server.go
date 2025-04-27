package server

import (
	"net/http"

	"github.com/gogineni1998/oolio-assignment-backend/authentication"
	"github.com/gogineni1998/oolio-assignment-backend/handlers"
)

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/token", handlers.Token())
	mux.HandleFunc("/register", handlers.Register())
	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		authentication.ValidateToken(handlers.Products()).ServeHTTP(w, r)
	})
	mux.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		authentication.ValidateToken(handlers.Product()).ServeHTTP(w, r)
	})
	mux.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		authentication.ValidateToken(handlers.Order()).ServeHTTP(w, r)
	})
	return mux
}
