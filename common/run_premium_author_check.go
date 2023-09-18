package common

import (
	db "author-handler-service/db"
	"author-handler-service/lib"
	"context"
	"log"
)

func RunPremiumAuthorCheck(config *lib.Config, ctx context.Context) {
	authors := db.GetAllAuthors(config, ctx)
	authorEligibilityConfig := db.GetAuthorEligiblityConfig(config, ctx)
	for _, author := range authors {
		isPremium := GetIsPremiumAuthorCondition(author, authorEligibilityConfig)
		err := UpdateAuthorPremiumStatus(config, author.AuthorId, isPremium)
		if err != nil {
			log.Printf("Error updating author: %v", err)
		}
	}
}
