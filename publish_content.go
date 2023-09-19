package main

import (
	commonUtils "author-handler-service/common"
	queueUtils "author-handler-service/common/queue_utils"
	db "author-handler-service/db"
	authorModel "author-handler-service/models/author_details"
	queueModel "author-handler-service/models/queue"
	"time"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PublishContent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	//This API will allow an author to post content & sync isPremiumAuthor condition realtime using events.

	var payload authorModel.AddAuthorContent
	err := json.NewDecoder(r.Body).Decode(&payload)
	//Validating payload against author struct

	log.Printf("Payload received inside publish content %v with request ID %v", payload, ctx.Value("REQUEST_ID"))

	if err != nil {
		log.Printf("%d: Invalid input %s", ctx, err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	authorDoc, _ := db.UpdateAuthorDetails(&config, payload)
	log.Printf("Author with ID %s isPremiumAuthor %v with request ID %v", payload.AuthorId, authorDoc.IsPremiumAuthor, ctx.Value("REQUEST_ID"))
	log.Printf("Updated Author doc: %v", authorDoc)
	if authorDoc.IsPremiumAuthor {
		//Author is already a premium author & content will be published to premium users.
		commonUtils.MockSendContentToPremiumUsers()
		log.Printf("Successfully published content to premium users for author %s with request ID %v", authorDoc.AuthorId, ctx.Value("REQUEST_ID"))
	} else {
		//Author is not premium author yet hence triggering message in queue to check for
		//premium author condition & update realtime if author has now become premium author.
		log.Printf("Author is not a premium author yet hence sending authorId %v to queue with request ID %v", payload.AuthorId, ctx.Value("REQUEST_ID"))
		queueName := config.QueueName
		producerPayload := queueModel.NewQueueStruct(payload.AuthorId, queueName)
		log.Printf("Sending payload to producer %v with request ID %v", producerPayload, ctx.Value("REQUEST_ID"))
		queueUtils.Producer(&config, producerPayload, ctx)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Content published successfully")
}
