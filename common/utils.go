package common

import (
	db "author-handler-service/db"
	"author-handler-service/lib"
	authorModel "author-handler-service/models/author_details"
	authorEligiblityConfig "author-handler-service/models/author_eligiblity_config"
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

func UpdateAuthorPremiumStatus(config *lib.Config, authorID string, isPremium bool) error {
	updateAuthorStruct := db.NewAuthorPremiumStruct(isPremium, authorID)
	return db.UpdateAuthorPremiumStatus(config, updateAuthorStruct)
}
