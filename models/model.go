package models

import "github.com/golang-jwt/jwt/v5"

var (
	Users = map[string]string{}
)

type ContextKey string

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Product struct {
	ID       string  `bson:"id" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Category string  `bson:"category" json:"category"`
}

type Item struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type OrderPayload struct {
	CouponCode string `json:"couponCode"`
	Items      []Item `json:"items"`
}

type Order struct {
	Username string     `json:"username"`
	Items    []Item     `json:"items"`
	Products []*Product `json:"products"`
}

type OrderResponse struct {
	ID       any        `json:"id"`
	Items    []Item     `json:"items"`
	Products []*Product `json:"products"`
}
