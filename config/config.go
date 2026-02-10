package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

// LoadConfig loads environment variables from .env
func LoadConfig() *Config {
	// Load .env file (optional, won't fail if not found)
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	// Read PORT from environment
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not set, defaulting to 8080")
		port = "8080"
	}

	return &Config{
		Port: port,
	}
}
