package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gogineni1998/oolio-assignment-backend/configuration"
	"github.com/gogineni1998/oolio-assignment-backend/database"
)

var Product = func() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Invalid request type", http.StatusForbidden)
			return
		}
		path := strings.TrimPrefix(r.URL.Path, "/product/")
		productID := strings.TrimSuffix(path, "/")

		if productID == "" {
			http.Error(w, "Invalid ID supplied", http.StatusBadRequest)
			return
		}
		databaseProduct, err := database.GetProductByID(configuration.DBProductsCollection, productID)
		if err != nil {
			http.Error(w, "Error fetching product", http.StatusInternalServerError)
			return
		}
		if databaseProduct == nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(databaseProduct); err != nil {
			http.Error(w, "Error encoding product to JSON", http.StatusInternalServerError)
			return
		}
	}
}

var Products = func() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Invalid request type", http.StatusForbidden)
			return
		}
		products, err := database.GetAllProducts(configuration.DBProductsCollection)
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "Error encoding products to JSON", http.StatusInternalServerError)
			return
		}
	}
}
