package main

import (
	db "author-handler-service/db"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	commonUtils "author-handler-service/common"

	"github.com/gorilla/mux"
)

// API to simulate author follower count
func SimulateAuthorFollowers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	//Geneartes random number between +100 to -100 and updates noOfFollowers in author table.
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(201) - 100
	log.Printf("Changing authorId %s followers by %v with request ID %v", authorId, randomNumber, ctx.Value("REQUEST_ID"))
	aurthorFollowerCountPayload := db.NewAuthorFollowerCountStruct(authorId, randomNumber)
	updatedAuthorDoc, _ := db.UpdateAuthorFollowerCount(&config, aurthorFollowerCountPayload)
	log.Printf("Simulated follower count for author %v : %v with request ID %v", authorId, updatedAuthorDoc, ctx.Value("REQUEST_ID"))
	//Sync isPremiumAuthor condition & update the flag based on followers going up / down
	go commonUtils.ComputeAndUpdateIsPremiumAuthor(&config, authorId, ctx)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Simulated author's follower counts successfully!")
}
