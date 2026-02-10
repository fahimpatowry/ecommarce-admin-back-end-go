package main

// @title Initial Project API
// @version 1.0
// @description REST API with Go, MongoDB, and Swagger
// @contact.email fahim@example.com

import (
	"log"
	"os"

	"initial_project/config"
	"initial_project/database"
	"initial_project/server"

	"github.com/joho/godotenv"

	_ "initial_project/docs" 
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration (e.g., PORT)
	cfg := config.LoadConfig()

	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")
	db, err := database.ConnectMongo(mongoURI, mongoDB)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Insert a sample user
	database.InsertSampleUser(db)

	// Initialize server
	s := server.NewServer(db)

	// Add routes


	s.Routes()

	// Start server
	log.Printf("Server running on port %s\n", cfg.Port)
	if err := s.Start(cfg.Port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
