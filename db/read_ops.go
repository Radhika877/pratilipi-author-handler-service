package db

import (
	"author-handler-service/lib"
	authorModel "author-handler-service/models/author_details"
	eligiblityConfigModel "author-handler-service/models/author_eligiblity_config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllAuthors(config *lib.Config, ctx context.Context) []authorModel.AuthorDetails {
	mongoClient := config.MongoDBCollection.MongoClient
	collection := mongoClient.Database("author-db").Collection("author-details")
	var result []authorModel.AuthorDetails
	curr, err := collection.Find(ctx, bson.M{"isAuthor": true})
	if err != nil {
		log.Println(err)
	}
	for curr.Next(ctx) {
		var authorData authorModel.AuthorDetails
		if err := curr.Decode(authorData); err != nil {
			log.Printf("Error decoding author data %s", err)
		}
		result = append(result, authorData)
	}
	return result
}

func GetAuthorByAuthorId(config *lib.Config, authorId string) authorModel.AuthorDetails {
	var result authorModel.AuthorDetails
	mongoClient := config.MongoDBCollection.MongoClient
	collection := mongoClient.Database("author-db").Collection("author-details")
	filter := bson.M{"authorId": authorId}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Printf("Error while fetching author by id %s request ID %v", err, context.Background())
	}
	return result
}

func GetAuthorEligiblityConfig(config *lib.Config, ctx context.Context) eligiblityConfigModel.AuthorEligiblityConfig {
	mongoClient := config.MongoDBCollection.MongoClient
	collection := mongoClient.Database("author-db").Collection("author-eligiblity-config")
	var result eligiblityConfigModel.AuthorEligiblityConfig
	docID, _ := primitive.ObjectIDFromHex(config.ConfigDocID)
	filter := bson.M{"_id": docID}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Printf("Error while fetching author eligblity config from db %s", err)
	}
	return result
}
