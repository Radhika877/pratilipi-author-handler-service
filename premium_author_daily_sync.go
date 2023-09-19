package main

import (
	commonUtils "author-handler-service/common"
	"context"
	"fmt"
	"net/http"
	"time"
)

func PremiumAuthorDailySync(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "REQUEST_ID", time.Now().UnixNano())
	go commonUtils.RunPremiumAuthorCheck(&config, ctx)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Completed daily sync to update premium authors")

}
