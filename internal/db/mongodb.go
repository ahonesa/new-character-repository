package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ahonesa/new-character-repository/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBClient is a global MongoDB client accessible from anywhere in the application.
var MongoDBClient *mongo.Client

// ConnectToMongoDB initializes the MongoDB connection.
func ConnectToMongoDB(cfg config.AppConfig) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use the MongoURI from the AppConfig struct.
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection is successful.
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB successfully: " + cfg.MongoURI)
	MongoDBClient = client
}

// GetMongoDBClient returns the global MongoDB client.
func GetMongoDBClient() *mongo.Client {
	return MongoDBClient
}
