package configuration

import "github.com/gogineni1998/oolio-assignment-backend/models"

const (
	Address            string = "0.0.0.0:8080"
	UsernameContextKey        = models.ContextKey("username")
)

var (
	JwtKey = []byte("dGp8mVzXeWvYhRJflA0LZcwPUujqX6TbyFQ1KOHgBMsiEnrNYk")
)

func init() {
	// Initialization logic can go here if needed, but addr is already set.
}
