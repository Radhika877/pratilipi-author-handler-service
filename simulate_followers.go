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
	authorId := vars["authorID"]
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(201) - 100
	aurthorFollowerCountPayload := db.NewAuthorFollowerCountStruct(authorId, randomNumber)
	updatedAuthorDoc, _ := db.UpdateAuthorFollowerCount(&config, aurthorFollowerCountPayload)
	log.Printf("Simulated follower count for author %v : %v", authorId, updatedAuthorDoc)
	go commonUtils.ComputeAndUpdateIsPremiumAuthor(&config, authorId)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Simulated author's follower counts successfully!")
}
