package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo connects to MongoDB and returns a *mongo.Database
func ConnectMongo(uri, dbName string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping to make sure connection is alive
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database(dbName), nil
}

// InsertSampleUser inserts a test user into the "users" collection
func InsertSampleUser(db *mongo.Database) error {
	collection := db.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := bson.D{
		{"name", "Fahim"},
		{"role", "Developer"},
		{"experience", 3.4},
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	fmt.Println("Inserted document ID:", result.InsertedID)
	return nil
}
