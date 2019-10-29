package usecase

import (
	"context"
	"log"

	"github.com/dappface/news-loader/domain/model"
	"github.com/dappface/news-loader/domain/service"
)

var links = []string{
	"https://feeds.feedburner.com/CoinDesk",
	"https://cointelegraph.com/rss",
}

func NewRSS(r service.RSS) RSS {
	return &rss{r}
}

type RSS interface {
	LoadAll()
}

type rss struct {
	rssService service.RSS
}

func (s *rss) LoadAll() {
	log.Println("Start loading RSS entries")

	ctx := context.Background()
	ps := make([]model.RSSEntryPost, 0)
	for _, l := range links {
		log.Println("Fetching:", l)
		p, err := s.rssService.Parse(ctx, l)
		if err != nil {
			log.Println(err)
		}
		ps = append(ps, *p...)
	}

	if err := s.rssService.SavePosts(ctx, ps); err != nil {
		log.Println(err)
	}

	log.Println("Finished loading RSS entries")
}
