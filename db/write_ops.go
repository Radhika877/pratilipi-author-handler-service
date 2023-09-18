package db

import (
	"author-handler-service/lib"
	authorModel "author-handler-service/models/author_details"
	authoreligiblityconfig "author-handler-service/models/author_eligiblity_config"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthorPremiumStruct struct {
	IsPremiumAuthor bool
	AuthorId        string
}

func NewAuthorPremiumStruct(isPremiumAuthor bool, authorId string) AuthorPremiumStruct {
	return AuthorPremiumStruct{
		IsPremiumAuthor: isPremiumAuthor,
		AuthorId:        authorId,
	}
}

func UpdateAuthorPremiumStatus(config *lib.Config, authorPayload AuthorPremiumStruct) error {
	mongoClient := config.MongoDBCollection.MongoClient
	collection := mongoClient.Database("author-db").Collection("author-details")
	filter := bson.M{"authorId": authorPayload.AuthorId}
	update := bson.M{
		"$set": bson.M{"isPremiumAuthor": authorPayload.IsPremiumAuthor},
	}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	var updatedDocument authorModel.AuthorDetails
	res := collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedDocument)
	if res != nil {
		fmt.Printf("Error: %v\n", res)
		return res
	} else {
		fmt.Printf("Updated document: %+v\n", updatedDocument)
		return nil
	}
}

func UpdateAuthorDetails(config *lib.Config, payload authorModel.AddAuthorContent) (authorModel.AuthorDetails, error) {
	mongoClient := config.MongoDBCollection.MongoClient
	collection := mongoClient.Database("author-db").Collection("author-details")
	filter := bson.M{"authorId": payload.AuthorId}
	update := bson.M{
		"$inc": bson.M{"noOfPosts": 1},
		"$set": bson.M{"lastPublishedAt": time.Now()},
	}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedDocument authorModel.AuthorDetails
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	err := collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedDocument)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return updatedDocument, err
	}
	fmt.Printf("Updated document: %v", updatedDocument)
	return updatedDocument, nil
}

func UpdateAuthorEligiblityConfig(config *lib.Config, authorEligiblityConfig authoreligiblityconfig.AuthorEligiblityConfig) error {
	mongoClient := config.MongoDBCollection.MongoClient
	collection := mongoClient.Database("author-db").Collection("author-eligiblity-config")
	docID, _ := primitive.ObjectIDFromHex(config.ConfigDocID)
	filter := bson.M{"_id": docID}
	updateParams := authorEligiblityConfig.UpdateAuthorEligiblityConfigParams()
	update := bson.M{"$set": updateParams}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	var updatedDocument authorModel.AuthorDetails
	res := collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedDocument)
	if res != nil {
		fmt.Printf("Error: %v\n", res)
		return res
	} else {
		fmt.Printf("Updated document: %+v\n", updatedDocument)
		return nil
	}
}
