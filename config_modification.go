package main

import (
	"author-handler-service/db"
	authoreligiblityconfig "author-handler-service/models/author_eligiblity_config"
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
	var payload authoreligiblityconfig.AuthorEligiblityConfig
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("%d: Invalid input, %s", ctx.Value("REQUEST_ID"), err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	updateAuthorConfigErr := db.UpdateAuthorEligiblityConfig(&config, payload)
	if updateAuthorConfigErr != nil {
		log.Printf("Error in updating author eligiblity config %s", updateAuthorConfigErr)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated author eligiblity config successfully.")
}
