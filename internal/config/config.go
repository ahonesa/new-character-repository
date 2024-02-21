package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig holds the configuration values for the application.
type AppConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	MongoURI           string
	CookieKey          string
}

// LoadConfig reads configuration from environment variables.
func LoadConfig(envFile string) AppConfig {
	// Load .env file if it exists
	if err := godotenv.Load(envFile); err != nil {
		log.Println("No .env file found")
	}

	return AppConfig{
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		MongoURI:           getEnv("MONGO_URI", ""),
		CookieKey:          getEnv("COOKIE_KEY", ""),
	}
}

// getEnv reads an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
