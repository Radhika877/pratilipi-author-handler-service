package main

import (
	commonUtils "author-handler-service/common"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func PremiumAuthorDailySync(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	log.Printf("Starting Daily sync with request ID %v", ctx.Value("REQUEST_ID"))
	go commonUtils.RunPremiumAuthorCheck(&config, ctx)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Completed daily sync to update premium authors")

}
