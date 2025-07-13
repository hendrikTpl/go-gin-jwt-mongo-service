package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var dbName string

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := getEnv("DB_HOST", "mongodb")
	dbPort := getEnv("DB_PORT", "27017")
	dbUser := getEnv("DB_USER", "admin")
	dbPass := getEnv("DB_PASSWORD", "admin")
	dbName = getEnv("DB_NAME", "ht-go-db")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin", dbUser, dbPass, dbHost, dbPort)
	fmt.Println("Connecting to MongoDB:", uri)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	// ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("Connected to MongoDB")
	mongoClient = client
}

func GetCollection(name string) *mongo.Collection {
	if mongoClient == nil {
		log.Fatal("MongoDB client not initialized. Call ConnectDB() first.")
	}
	return mongoClient.Database(dbName).Collection(name)
}
