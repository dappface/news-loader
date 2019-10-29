package usecase

import (
	"context"
	"log"

	"github.com/dappface/news-loader/domain/service"
)

var (
	screenName = "dappface_com"
	lists      = []string{
		"devcon4-individuals",
		"web3summit-groups",
		"web3summit-individuals",
	}
)

func NewTwitter(s service.Twitter) Twitter {
	return &twitterU{s}
}

type Twitter interface {
	LoadAll()
}

type twitterU struct {
	s service.Twitter
}

func (u twitterU) LoadAll() {
	log.Println("Start loading tweets")

	ctx := context.Background()
	ts, lts, err := u.s.GetListTweetsBySlugs(ctx, lists, screenName)
	if err != nil {
		log.Println(err)
	}

	ps := u.s.CvtTweetsToPosts(*ts)
	filtered, err := u.s.FilterUniquePosts(ctx, ps)
	if err != nil {
		log.Println(err)
	}

	if len(*filtered) == 0 {
		log.Println("No new tweet to add")
		return
	}

	err = u.s.SavePosts(ctx, *filtered, *lts)
	if err != nil {
		log.Println(err)
	}

	log.Println("Finished loading tweets")
}
