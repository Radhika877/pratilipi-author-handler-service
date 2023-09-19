package common

import (
	db "author-handler-service/db"
	"author-handler-service/lib"
	"context"
	"log"
)

func RunPremiumAuthorCheck(config *lib.Config, ctx context.Context) bool {
	log.Printf("Started daily isPremiumAuthor sync with request ID %v", ctx.Value("REQUEST_ID"))
	authors, err := db.GetAllAuthors(config, ctx)
	if err != nil {
		log.Printf("Error in fetching all authors %s with request ID %v", err, ctx.Value("REQUEST_ID"))
	}
	log.Printf("Total no. of authors %v with request ID %v", len(authors), ctx.Value("REQUEST_ID"))
	isPremiumAuthorEligiblityConfig := db.GetAuthorEligiblityConfig(config, ctx)
	//Chunk authors if the size of author increases eventually.
	for _, author := range authors {
		authorDoc := db.GetAuthorByAuthorId(config, author.AuthorId)
		isPremiumAuthor := GetIsPremiumAuthorCondition(authorDoc, isPremiumAuthorEligiblityConfig)
		payload := db.NewAuthorPremiumStruct(isPremiumAuthor, author.AuthorId)
		db.UpdateAuthorPremiumStatus(config, payload)
	}
	log.Printf("Finished daily isPremiumAuthor sync with request ID %v", ctx.Value("REQUEST_ID"))
	return true
}
