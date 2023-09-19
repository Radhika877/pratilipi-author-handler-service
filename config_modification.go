package main

import (
	"author-handler-service/db"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ConfigModification(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	var payload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("%d: Invalid input, %s", ctx.Value("REQUEST_ID"), err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	log.Printf("Payload for modifying global config %v with request ID %v", payload, ctx.Value("REQUEST_ID"))
	updateAuthorConfigErr := db.UpdateAuthorEligibilityConfig(&config, payload)
	if updateAuthorConfigErr != nil {
		log.Printf("Error in updating author eligiblity config %s with request ID %v", updateAuthorConfigErr, ctx.Value("REQUEST_ID"))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated author eligiblity config successfully.")
}
