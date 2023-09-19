package common

import (
	db "author-handler-service/db"
	"author-handler-service/lib"
	"context"
	"log"
)

func RunPremiumAuthorCheck(config *lib.Config, ctx context.Context) {
	authors := db.GetAllAuthors(config, ctx)
	for _, author := range authors {
		err := ComputeAndUpdateIsPremiumAuthor(config, author.AuthorId)
		if err != nil {
			log.Printf("Error updating author: %v", err)
		}
	}
}
