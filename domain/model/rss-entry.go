package model

import "time"

type RSSEntry struct {
	Author       string    `firestore:"author"`
	Categories   []string  `firestore:"categories"`
	Description  string    `firestore:"description"`
	FeedTitle    string    `firestore:"feedTitle"`
	GUID         string    `firestore:"guid"`
	ImageURL     string    `firestore:"imageUrl"`
	Language     string    `firestore:"language"`
	PublisherURL string    `firestore:"publisherUrl"`
	Title        string    `firestore:"title"`
	URL          string    `firestore:"url"`
	CreatedAt    time.Time `firestore:"createdAt"`
}
