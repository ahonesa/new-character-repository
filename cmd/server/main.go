package main

import (
	"log"
	"github.com/ahonesa/new-character-repository/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	// Example: Use cfg.MongoURI to connect to MongoDB
	// db.Connect(cfg.MongoURI)

	log.Printf("Starting server with Google Client ID: %s", cfg.GoogleClientID)
	// Continue with application setup...
}
