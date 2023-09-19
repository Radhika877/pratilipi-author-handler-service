package common

import (
	db "author-handler-service/db"
	"author-handler-service/lib"
	authorModel "author-handler-service/models/author_details"
	authorEligiblityConfig "author-handler-service/models/author_eligiblity_config"
	"context"
	"log"
	"time"
)

// Returns if an author is eligible to become a premium author by comparing against the global configuration
func GetIsPremiumAuthorCondition(author authorModel.AuthorDetails, authorEligiblityConfig authorEligiblityConfig.AuthorEligiblityConfig) bool {
	if author.NoOfFollowers >= authorEligiblityConfig.MinimumNoOfFollowers {
		nDaysAgo := time.Now().AddDate(0, 0, -authorEligiblityConfig.WindowInDays)
		if author.NoOfPosts >= authorEligiblityConfig.MinimumNoOfPosts && author.LastPublishedAt.After(nDaysAgo) {
			return true
		}
	}
	return false
}

// This is a dummy func to trigger notifications to premium subscription users
// when a premium author posts content
func MockSendContentToPremiumUsers() bool {
	return true
}

func ComputeAndUpdateIsPremiumAuthor(config *lib.Config, authorId string, ctx context.Context) error {
	eligiblityConfig := db.GetAuthorEligiblityConfig(config, ctx)
	authorDoc := db.GetAuthorByAuthorId(config, authorId)
	isPremiumAuthor := GetIsPremiumAuthorCondition(authorDoc, eligiblityConfig)
	log.Printf("isPremiumAuthor %v for aurhor %v with request ID %v", isPremiumAuthor, authorId, ctx.Value("REQUEST_ID"))
	if authorDoc.IsPremiumAuthor == isPremiumAuthor {
		log.Printf("isPremiumAuthor value is same in db hence not updating for author %v", authorId)
	}
	updateAuthorStruct := db.NewAuthorPremiumStruct(isPremiumAuthor, authorId)
	_, updatePremiumErr := db.UpdateAuthorPremiumStatus(config, updateAuthorStruct)
	if updatePremiumErr != nil {
		return updatePremiumErr
	}
	return nil
}
