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

// ConnectDB must be called from main.go after env is loaded
func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set in environment variables")
	}

	fmt.Println("Connecting to MongoDB:", uri)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	// ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("✅ Connected to MongoDB")
	mongoClient = client
}

func GetCollection(name string) *mongo.Collection {
	if mongoClient == nil {
		log.Fatal("MongoDB client not initialized. Call ConnectDB() first.")
	}
	return mongoClient.Database("ginjwt").Collection(name)
}
