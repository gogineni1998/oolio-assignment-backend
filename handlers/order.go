package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gogineni1998/oolio-assignment-backend/configuration"
	"github.com/gogineni1998/oolio-assignment-backend/database"
	"github.com/gogineni1998/oolio-assignment-backend/models"
)

var Order = func() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request type", http.StatusForbidden)
			return
		}

		username := r.Context().Value(configuration.UsernameContextKey).(string)
		var payload models.OrderPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			log.Println("Error decoding payload:", err)
			return
		}
		defer r.Body.Close()
		if payload.CouponCode != "" && len(payload.CouponCode) < 8 && len(payload.CouponCode) > 10 {
			http.Error(w, "Invalid coupon code", http.StatusBadRequest)
			log.Println("Invalid coupon code:", payload.CouponCode)
			return
		}
		validateCoupon := 0
		for _, coupon := range configuration.CopounData {
			if strings.Contains(coupon, payload.CouponCode) {
				validateCoupon = validateCoupon + 1
				if validateCoupon == 2 {
					break
				}
			}
		}
		if validateCoupon < 2 {
			http.Error(w, "Invalid coupon code", http.StatusBadRequest)
			log.Println("Invalid coupon code:", payload.CouponCode)
			return
		}
		productIds := []string{}
		for _, product := range payload.Items {
			productIds = append(productIds, product.ProductID)
		}
		products, err := database.GetProductsByIDs(configuration.DBProductsCollection, productIds)
		if err != nil {
			http.Error(w, "Error fetching products", http.StatusInternalServerError)
			log.Println("Error fetching products:", err)
		}
		orderID, err := database.InsertOrder(configuration.DBOrdersCollection, models.Order{
			Items:    payload.Items,
			Products: products,
			Username: username,
		})
		if err != nil {
			http.Error(w, "Error creating order", http.StatusInternalServerError)
			log.Println("Error creating order:", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := models.OrderResponse{
			ID:       orderID,
			Items:    payload.Items,
			Products: products,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response to JSON", http.StatusInternalServerError)
			log.Println("Error encoding response:", err)
			return
		}
	}

}
