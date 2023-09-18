package lib

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBCollection struct {
	MongoClient *mongo.Client
}

// Initialise and return mongoclient
func InitialiseMongoDb(config *Config) mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.MongodbURI).SetServerAPIOptions(serverAPI)
	mongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("couldn't create Mongo client")
	}

	fmt.Println("Successfully connected to MongoDB!")
	config.MongoDBCollection = &MongoDBCollection{
		MongoClient: mongoClient,
	}
	return *mongoClient
}
