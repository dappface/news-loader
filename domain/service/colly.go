package service

import (
	"log"

	"github.com/gocolly/colly"
)

func NewColly() *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)
	c.OnError(errorHandler)

	return c
}

func errorHandler(r *colly.Response, err error) {
	log.Printf("Request URL: %v failed with response: %v \nError %v", r.Request.URL, r.StatusCode, err)
}
