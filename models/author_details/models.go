package authordetails

import "time"

type AuthorDetails struct {
	AuthorId        string    `json:"authorID"`
	AuthorName      string    `json:"authorName,omitempty"`
	CreatedAt       time.Time `bson:"createdAt"`
	NoOfFollowers   int       `json:"noOfFollowers,omitempty"`
	LastPublishedAt time.Time `bson:"lastPublishedAt"`
	NoOfPosts       int       `json:"noOfPosts,omitempty"`
	IsPremiumAuthor bool      `json:"isPremiumAuthor"`
	IsPremiumUser   bool      `json:"isPremiumUser"`
	IsAuthor        bool      `json:"isAuthor"`
}

type AddAuthorContent struct {
	AuthorId string `json:"authorID"`
	Content  string `json:"content"`
}

type ConsumerData struct {
	AuthorId string `json:"authorID"`
}
