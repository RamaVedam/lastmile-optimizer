package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Load .env files
)

var (
	DatabaseURL string
	APIPort     string
)

// LoadConfig initializes configuration from environment variables or .env file
func LoadConfig() {
	// Load from .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading environment variables...")
	}

	DatabaseURL = getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/lastmile_db?sslmode=disable")
	APIPort = getEnv("API_PORT", "8080")
}

// getEnv fetches the value of an environment variable or falls back to a default
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
