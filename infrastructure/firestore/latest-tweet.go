package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/dappface/news-loader/domain/model"
	"github.com/dappface/news-loader/domain/repository"
)

func NewLatestTweet() repository.LatestTweet {
	return &latestTweet{}
}

type latestTweet struct{}

func (s *latestTweet) GetSinceIDBySlug(ctx context.Context, sl string) (*string, error) {
	lts, err := Client.
		Collection(Collection.LatestTweets).
		Where("listName", "==", sl).
		OrderBy("createdAt", firestore.Desc).
		Limit(1).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}

	if len(lts) == 0 {
		return nil, nil
	}

	var lt model.LatestTweet
	lts[0].DataTo(&lt)

	return &lt.ID, nil
}
