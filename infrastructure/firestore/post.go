package firestore

import (
	"context"

	"github.com/dappface/news-loader/domain/model"
	"github.com/dappface/news-loader/domain/repository"
)

func NewPost() repository.Post {
	return &post{}
}

type post struct{}

func (r *post) AddRSSEntries(ctx context.Context, ps []model.RSSEntryPost) error {
	for _, p := range ps {
		col := Client.Collection(Collection.Posts)
		_, _, err := col.Add(ctx, p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *post) AddTweetsWithLatestTweets(ctx context.Context, ps []model.TweetPost, lts []model.LatestTweet) error {
	batch := Client.Batch()
	postsCol := Client.Collection(Collection.Posts)
	ltsCol := Client.Collection(Collection.LatestTweets)
	for _, p := range ps {
		batch.Create(postsCol.NewDoc(), p)
	}

	for _, lt := range lts {
		batch.Create(ltsCol.NewDoc(), lt)
	}

	_, err := batch.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *post) IsRSSEntryExist(ctx context.Context, p model.RSSEntryPost) (*bool, error) {
	col := Client.Collection(Collection.Posts)
	docSnap, err := col.Where("postDataId", "==", p.PostDataID).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	b := len(docSnap) == 1
	return &b, err
}

func (r *post) IsTweetExist(ctx context.Context, p model.TweetPost) (*bool, error) {
	col := Client.Collection(Collection.Posts)
	docSnap, err := col.Where("postDataId", "==", p.PostDataID).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	b := len(docSnap) == 1
	return &b, err
}
