package repository

import (
	"context"

	"github.com/dappface/news-loader/domain/model"
)

type Post interface {
	AddRSSEntries(ctx context.Context, posts []model.RSSEntryPost) error
	AddTweetsWithLatestTweets(ctx context.Context, posts []model.TweetPost, latestTweets []model.LatestTweet) error
	IsRSSEntryExist(ctx context.Context, posts model.RSSEntryPost) (*bool, error)
	IsTweetExist(ctx context.Context, posts model.TweetPost) (*bool, error)
}
