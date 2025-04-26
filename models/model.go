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
