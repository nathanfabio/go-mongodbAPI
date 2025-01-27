package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// ConnMongo connects to MongoDB
func ConnMongo() (*mongo.Client, error) {
	// connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// getting username and password from environment variables
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	// setting auth
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Connected to MongoDB!")

	return client, nil
}

// GetCollection returns the collection
func GetCollection() *mongo.Collection {
	return collection
}