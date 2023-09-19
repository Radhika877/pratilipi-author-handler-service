package db

import (
	"author-handler-service/lib"
	authorModel "author-handler-service/models/author_details"
	"context"
	"log"
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

type AuthorFollowerCountStruct struct {
	AuthorId      string
	NoOfFollowers int
}

func NewAuthorFollowerCountStruct(authorId string, noOfFollowers int) AuthorFollowerCountStruct {
	return AuthorFollowerCountStruct{
		AuthorId:      authorId,
		NoOfFollowers: noOfFollowers,
	}
}

func UpdateAuthorPremiumStatus(config *lib.Config, authorPayload AuthorPremiumStruct) (authorModel.AuthorDetails, error) {
	filter := bson.M{"authorId": authorPayload.AuthorId}
	update := bson.M{
		"$set": bson.M{"isPremiumAuthor": authorPayload.IsPremiumAuthor},
	}
	updatedDoc, updateErr := findOneAndUpdate(config, "author-details", authorPayload.AuthorId, filter, update)
	if updateErr != nil {
		return updatedDoc, updateErr
	}
	return updatedDoc, nil
}

func UpdateAuthorDetails(config *lib.Config, payload authorModel.AddAuthorContent) (authorModel.AuthorDetails, error) {
	filter := bson.M{"authorId": payload.AuthorId}
	update := bson.M{
		"$inc": bson.M{"noOfPosts": 1},
		"$set": bson.M{"lastPublishedAt": primitive.Timestamp{T: uint32(time.Now().Unix()), I: 0}},
	}
	updatedDoc, updateErr := findOneAndUpdate(config, "author-details", payload.AuthorId, filter, update)
	if updateErr != nil {
		return updatedDoc, updateErr
	}
	return updatedDoc, nil
}

func UpdateAuthorFollowerCount(config *lib.Config, authorFollowerCount AuthorFollowerCountStruct) (authorModel.AuthorDetails, error) {
	filter := bson.M{"authorId": authorFollowerCount.AuthorId}
	update := bson.M{
		"$inc": bson.M{"noOfFollowers": authorFollowerCount.NoOfFollowers},
	}
	updatedDoc, updateErr := findOneAndUpdate(config, "author-details", authorFollowerCount.AuthorId, filter, update)
	if updateErr != nil {
		return updatedDoc, updateErr
	}
	return updatedDoc, nil
}

func UpdateAuthorEligibilityConfig(config *lib.Config, updates map[string]interface{}) error {
	mongoClient := config.MongoDBCollection.MongoClient
	collection := mongoClient.Database("author-db").Collection("author-eligiblity-config")
	docID, _ := primitive.ObjectIDFromHex(config.ConfigDocID)
	filter := bson.M{"_id": docID}
	update := bson.M{"$set": updates}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	var updatedDocument authorModel.AuthorDetails
	res := collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedDocument)
	if res != nil {
		log.Printf("Error: %v", res)
		return res
	} else {
		log.Printf("Updated document: %+v", updatedDocument)
		return nil
	}
}

func findOneAndUpdate(config *lib.Config, collectionName string, authorId string, filter primitive.M, update primitive.M) (authorModel.AuthorDetails, error) {
	mongoClient := config.MongoDBCollection.MongoClient
	collection := mongoClient.Database("author-db").Collection(collectionName)
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updatedDocument authorModel.AuthorDetails
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	err := collection.FindOneAndUpdate(ctx, filter, update, options).Decode(&updatedDocument)
	if err != nil {
		log.Printf("Error: %s", err)
		return updatedDocument, err
	}
	log.Printf("Updated document: %v", updatedDocument)
	return updatedDocument, nil
}
