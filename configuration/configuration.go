package configuration

import (
	"bufio"
	"context"
	"os"

	"log"

	"github.com/gogineni1998/oolio-assignment-backend/coupons"
	"github.com/gogineni1998/oolio-assignment-backend/database"
	"github.com/gogineni1998/oolio-assignment-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Address              string = "0.0.0.0:5000"
	UsernameContextKey          = models.ContextKey("username")
	JwtKey                      = []byte("dGp8mVzXeWvYhRJflA0LZcwPUujqX6TbyFQ1KOHgBMsiEnrNYk")
	couponsFilePaths     []string
	CopounData           = make(map[string]string)
	DBClient             *mongo.Client
	DBProductsCollection *mongo.Collection
	DBOrdersCollection   *mongo.Collection
	DBUsersCollection    *mongo.Collection
)

func init() {
	DBClient = database.DatabaseConnect()
	if DBClient == nil {
		panic("Failed to connect to the database")
	}
	DBProductsCollection = database.ConnectMongoCollection(DBClient, "oolio-product-database", "products")
	DBOrdersCollection = database.ConnectMongoCollection(DBClient, "oolio-product-database", "orders")
	DBUsersCollection = database.ConnectMongoCollection(DBClient, "oolio-product-database", "users")
	ctx := context.Background()
	database.InsertDataIfNotExists(DBProductsCollection, ctx)
	LoadCouponsFilePaths("./local_configuration")
	for _, filePath := range couponsFilePaths {
		CopounData[filePath] = coupons.LoadCoupons(filePath)
	}
	log.Println("Coupons loaded successfully")
}

func LoadCouponsFilePaths(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open coupons file paths config: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			couponsFilePaths = append(couponsFilePaths, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed reading coupons file paths config: %v", err)
	}
}
