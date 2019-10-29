package service

import (
	"context"
	"log"
	"time"

	"github.com/dappface/news-loader/domain/model"
	"github.com/dappface/news-loader/domain/repository"
	"github.com/dappface/news-loader/infrastructure/twitter"
)

func NewTwitter(postRepository repository.Post, latestTweetRepository repository.LatestTweet) Twitter {
	return &twitterS{
		postRepository,
		latestTweetRepository,
	}
}

type Twitter interface {
	GetListTweetsBySlugs(ctx context.Context, slugs []string, ownerScreenName string) (*[]model.Tweet, *[]model.LatestTweet, error)
	CvtTweetsToPosts(tweets []model.Tweet) []model.TweetPost
	FilterUniquePosts(ctx context.Context, posts []model.TweetPost) (*[]model.TweetPost, error)
	SavePosts(ctx context.Context, posts []model.TweetPost, latestTweets []model.LatestTweet) error
}

type twitterS struct {
	postRepository        repository.Post
	latestTweetRepository repository.LatestTweet
}

func (s *twitterS) GetListTweetsBySlugs(ctx context.Context, sls []string, sn string) (*[]model.Tweet, *[]model.LatestTweet, error) {
	tweets := make([]model.Tweet, 0)
	lts := make([]model.LatestTweet, 0)
	for _, sl := range sls {
		log.Printf("Fetching: %+v", sl)

		sinceID, err := s.latestTweetRepository.GetSinceIDBySlug(ctx, sl)
		if err != nil {
			return nil, nil, err
		}

		if sinceID == nil {
			log.Println("Since ID:", sinceID)
		} else {
			log.Println("Since ID:", *sinceID)
		}

		ts, err := twitter.Client.GetTweetsBySlug(sl, sn, sinceID)
		if err != nil {
			return nil, nil, err
		}

		tweets = append(tweets, *ts...)
		if len(*ts) != 0 {
			t := (*ts)[0]
			lts = append(lts, model.LatestTweet{
				ID:        t.IDStr,
				ListName:  sl,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			})
		}
	}

	return &tweets, &lts, nil
}

func (s *twitterS) CvtTweetsToPosts(ts []model.Tweet) []model.TweetPost {
	posts := make([]model.TweetPost, 0)
	for _, t := range ts {
		posts = append(posts, model.TweetPost{
			PostDataID:  t.IDStr,
			PostType:    model.PostType.Tweet,
			PostData:    t,
			Language:    t.Lang,
			PublishedAt: t.CreatedAt,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}
	return posts
}

func (s *twitterS) FilterUniquePosts(ctx context.Context, ps []model.TweetPost) (*[]model.TweetPost, error) {
	m := make(map[string]model.TweetPost, 0)
	for _, p := range ps {
		isExist, err := s.postRepository.IsTweetExist(ctx, p)
		if err != nil {
			return nil, err
		}
		if *isExist {
			log.Println("Post already exists:", p.PostDataID)
			continue
		}
		m[p.PostDataID] = p
	}

	filtered := make([]model.TweetPost, 0)
	for _, v := range m {
		filtered = append(filtered, v)
	}

	return &filtered, nil
}

func (s *twitterS) SavePosts(ctx context.Context, ps []model.TweetPost, lts []model.LatestTweet) error {
	if err := s.postRepository.AddTweetsWithLatestTweets(ctx, ps, lts); err != nil {
		return err
	}
	return nil
}
