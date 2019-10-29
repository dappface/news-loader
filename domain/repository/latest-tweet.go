package repository

import "context"

type LatestTweet interface {
	GetSinceIDBySlug(ctx context.Context, slug string) (*string, error)
}
