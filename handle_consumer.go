package main

import (
	commonUtils "author-handler-service/common"
	"author-handler-service/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func HandleConsumer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())

	var payload json.Number
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("Invalid input %s with request ID %v", err, ctx.Value("REQUEST_ID"))
		return
	}
	log.Printf("Payload received in consumer handler %s request ID %v", payload, ctx.Value("REQUEST_ID"))
	author := db.GetAuthorByAuthorId(&config, payload.String())
	premiumAuthorEligibilityConfig := db.GetAuthorEligiblityConfig(&config, ctx)
	isPremiumAuthor := commonUtils.GetIsPremiumAuthorCondition(author, premiumAuthorEligibilityConfig)
	if author.IsPremiumAuthor == isPremiumAuthor {
		log.Printf("IsPremiumAuthor flag for author %s is the same hence not updating db with request ID %v", author.AuthorId, ctx.Value("REQUEST_ID"))
		fmt.Fprintf(w, "Status same hence not updating.")
	} else {
		log.Printf("isPremiumAuthor %v Condition for author %v with request ID %v", isPremiumAuthor, author.AuthorId, ctx.Value("REQUEST_ID"))
		updateData := db.NewAuthorPremiumStruct(isPremiumAuthor, author.AuthorId)
		_, updateStatusErr := db.UpdateAuthorPremiumStatus(&config, updateData)
		if updateStatusErr != nil {
			log.Printf("Error in updating premium author status: %s with request ID %v", updateStatusErr, ctx.Value("REQUEST_ID"))
			http.Error(w, "Error", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Update author premium status successfully")
	}
}
