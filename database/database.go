package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gogineni1998/oolio-assignment-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUsername = "goginenidheeraj"
	dbPassword = "Dheeraj123#"
	dbHost     = "oolio.pgcsmyl.mongodb.net"
	dbName     = "Oolio"
)

func DatabaseConnect() *mongo.Client {
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s", dbUsername, dbPassword, dbHost, dbName)

	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func DisconnectDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB!")
}

func ConnectMongoCollection(client *mongo.Client, databaseName, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

func InsertDataIfNotExists(collection *mongo.Collection, ctx context.Context) {
	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		fmt.Printf("Collection already has %d documents. Skipping insert.\n", count)
		return
	}

	file, err := os.Open("./data/data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var documents []any
	if err := json.Unmarshal(byteValue, &documents); err != nil {
		log.Fatal(err)
	}

	result, err := collection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %d documents!\n", len(result.InsertedIDs))
}

func GetAllProducts(collection *mongo.Collection) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductByID(collection *mongo.Collection, productID string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "id", Value: productID}}

	var product models.Product
	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No product found with ID: %s\n", productID)
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func GetProductsByIDs(collection *mongo.Collection, productIDs []string) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "id", Value: bson.D{{Key: "$in", Value: productIDs}}}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func InsertOrder(collection *mongo.Collection, order models.Order) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertResult, err := collection.InsertOne(ctx, order)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	orderId := insertResult.InsertedID
	return orderId, nil
}

func CheckUser(collection *mongo.Collection, username string) (*models.Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "username", Value: username}}

	var users models.Credentials
	err := collection.FindOne(ctx, filter).Decode(&users)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("No User found with ID: %s\n", username)
			return nil, nil
		}
		return nil, err
	}

	return &users, nil
}

func InsertUser(collection *mongo.Collection, credentials models.Credentials) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	insertResult, err := collection.InsertOne(ctx, credentials)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	userId := insertResult.InsertedID
	return userId, nil
}
