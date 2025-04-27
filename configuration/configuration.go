package configuration

import (
	"context"

	"github.com/gogineni1998/oolio-assignment-backend/coupons"
	"github.com/gogineni1998/oolio-assignment-backend/database"
	"github.com/gogineni1998/oolio-assignment-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Address            string = "0.0.0.0:8080"
	UsernameContextKey        = models.ContextKey("username")
)

var (
	JwtKey               = []byte("dGp8mVzXeWvYhRJflA0LZcwPUujqX6TbyFQ1KOHgBMsiEnrNYk")
	couponsFilePaths     = []string{"./data/couponbase1.gz", "./data/couponbase2.gz", "./data/couponbase3.gz"}
	CopounData           = make(map[string]string)
	DBClient             *mongo.Client
	DBProductsCollection *mongo.Collection
	DBOrdersCollection   *mongo.Collection
)

func init() {
	DBClient = database.DatabaseConnect()
	if DBClient == nil {
		panic("Failed to connect to the database")
	}
	DBProductsCollection = database.ConnectMongoCollection(DBClient, "oolio-product-database", "products")
	DBOrdersCollection = database.ConnectMongoCollection(DBClient, "oolio-product-database", "orders")
	ctx := context.Background()
	database.InsertDataIfNotExists(DBProductsCollection, ctx)
	for _, filePath := range couponsFilePaths {
		CopounData[filePath] = coupons.LoadCoupons(filePath)
	}
}
