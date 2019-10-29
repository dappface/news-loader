package service

import (
	"context"
	"log"
	"time"

	"github.com/dappface/news-loader/domain/model"
	"github.com/dappface/news-loader/domain/repository"
	"github.com/mmcdole/gofeed"
)

func NewRSS(r repository.Post) RSS {
	return &rssR{r}
}

type RSS interface {
	Parse(ctx context.Context, link string) (*[]model.RSSEntryPost, error)
	SavePosts(ctx context.Context, posts []model.RSSEntryPost) error
}

type rssR struct {
	r repository.Post
}

func (s *rssR) Parse(ctx context.Context, l string) (*[]model.RSSEntryPost, error) {
	ps := make([]model.RSSEntryPost, 0)
	fp := gofeed.NewParser()
	fp.RSSTranslator = NewTranslator()
	f, _ := fp.ParseURL(l)
	for _, e := range f.Items {
		pd := &model.RSSEntry{
			Author:       e.Author.Name,
			Categories:   e.Categories,
			Description:  e.Description,
			FeedTitle:    f.Title,
			GUID:         e.GUID,
			ImageURL:     e.Image.URL,
			Language:     f.Language,
			PublisherURL: f.Link,
			Title:        e.Title,
			URL:          e.Link,
			CreatedAt:    *e.PublishedParsed,
		}
		p := &model.RSSEntryPost{
			PostDataID:  e.GUID,
			PostType:    model.PostType.RSSEntry,
			PostData:    *pd,
			Language:    &f.Language,
			PublishedAt: *e.PublishedParsed,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		b, err := s.r.IsRSSEntryExist(ctx, *p)
		if err != nil {
			return nil, err
		}

		if *b {
			log.Printf("Post already exists: %v", p.PostDataID)
			continue
		}

		ps = append(ps, *p)
	}

	return &ps, nil
}

func (s *rssR) SavePosts(ctx context.Context, ps []model.RSSEntryPost) error {
	return s.r.AddRSSEntries(ctx, ps)
}
