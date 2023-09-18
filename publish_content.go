package main

import (
	commonUtils "author-handler-service/common"
	queueUtils "author-handler-service/common/queue_utils"
	db "author-handler-service/db"
	authorModel "author-handler-service/models/author_details"
	queueModel "author-handler-service/models/queue"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func PublishContent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	//This API will allow an author to post content & sync isPremiumAuthor condition realtime using events.

	var payload authorModel.AddAuthorContent
	err := json.NewDecoder(r.Body).Decode(&payload)
	//Validating payload against author struct

	log.Printf("Payload received inside publish content %v with request ID %v", payload, ctx)

	if err != nil {
		log.Printf("%d: Invalid input %s", ctx.Value("REQUEST_ID"), err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	authorDoc, _ := db.UpdateAuthorDetails(&config, payload)
	log.Printf("Author with ID %s isPremiumAuthor %v with Id %v", payload.AuthorId, authorDoc.IsPremiumAuthor, ctx)

	if authorDoc.IsPremiumAuthor {
		//Author is already a premium author & content will be published to premium users.
		commonUtils.MockSendContentToPremiumUsers()
		log.Printf("Successfully published content to premium users for author %s with request ID %v", authorDoc.AuthorId, ctx)
	} else {
		//Author is not premium author yet hence triggering message in queue to check for
		//premium author condition & update realtime if author has now become premium author.
		queueName := config.QueueName
		producerPayload := queueModel.NewQueueStruct(authorDoc.AuthorId, queueName)
		queueUtils.Producer(&config, producerPayload)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Content published successfully")
}
