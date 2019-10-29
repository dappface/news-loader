package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var (
	Client     *firestore.Client
	Collection = collection{
		Posts:        "posts",
		LatestTweets: "latestTweets",
	}
)

func init() {
	ctx := context.Background()
	var err error
	projectID := os.Getenv("PROJECT_ID")
	Client, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Failed to initialize firestore: %v", err)
	}
}

type collection struct {
	Posts        string
	LatestTweets string
}
