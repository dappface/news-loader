package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/mmcdole/gofeed"
	"github.com/mmcdole/gofeed/rss"
)

var imageURL string

func NewTranslator() *Translator {
	t := &Translator{}
	t.defaultTranslator = &gofeed.DefaultRSSTranslator{}
	return t
}

type Translator struct {
	defaultTranslator *gofeed.DefaultRSSTranslator
}

func (t *Translator) Translate(feed interface{}) (*gofeed.Feed, error) {
	rss, found := feed.(*rss.Feed)
	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *rss.Feed")
	}

	f, err := t.defaultTranslator.Translate(rss)
	if err != nil {
		return nil, err
	}

	switch rss.Link {
	case "https://www.coindesk.com":
		for _, e := range f.Items {
			fetchCoinDeskPostImageLink(e.Link)
			e.Image = &gofeed.Image{URL: imageURL, Title: ""}
		}
	case "https://cointelegraph.com":
		for _, e := range f.Items {
			desc, url := parseCointelegraphDescription(e.Description)
			e.Description = desc
			e.Image = &gofeed.Image{URL: url, Title: ""}
			author := parseCointelegraphAuthor(e.Author.Name)
			e.Author.Name = author
		}
	}

	return f, nil
}

func fetchCoinDeskPostImageLink(link string) {
	c := NewColly()
	c.OnHTML("article.coindesk-article", htmlHandler)
	c.Visit(link)
}

func htmlHandler(e *colly.HTMLElement) {
	url, exists := e.DOM.First().Find("div.coindesk-article-header-image source").Attr("srcset")
	if exists == false {
		log.Println("Failed to find ImageURL")
		imageURL = ""
	}
	imageURL = url
}

func parseCointelegraphDescription(s string) (string, string) {
	s = strings.Split(s, "<img src=\"")[1]
	sl := strings.Split(s, "\"><p>")
	d := strings.Split(sl[1], "</p>")[0]
	return d, sl[0]
}

func parseCointelegraphAuthor(s string) string {
	substr := " By "
	if strings.Contains(s, substr) {
		return strings.Split(s, substr)[1]
	}
	return s
}
