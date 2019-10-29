package model

import "time"

var (
	PostType = postType{
		RSSEntry: "rssEntry",
		Tweet:    "tweet",
	}
)

type RSSEntryPost struct {
	PostDataID  string    `firestore:"postDataId"`
	PostType    string    `firestore:"postType"` // "rssEntry" | "tweet"
	PostData    RSSEntry  `firestore:"postData"`
	Language    *string   `firestore:"language"`
	PublishedAt time.Time `firestore:"publishedAt"`
	CreatedAt   time.Time `firestore:"createdAt"`
	UpdatedAt   time.Time `firestore:"updatedAt"`
}

type TweetPost struct {
	PostDataID  string    `firestore:"postDataId"`
	PostType    string    `firestore:"postType"` // "rssEntry" | "tweet"
	PostData    Tweet     `firestore:"postData"`
	Language    *string   `firestore:"language"`
	PublishedAt time.Time `firestore:"publishedAt"`
	CreatedAt   time.Time `firestore:"createdAt"`
	UpdatedAt   time.Time `firestore:"updatedAt"`
}

type postType struct {
	RSSEntry string
	Tweet    string
}
