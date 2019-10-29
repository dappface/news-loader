package registry

import (
	"net/http"

	"github.com/dappface/news-loader/domain/service"
	"github.com/dappface/news-loader/infrastructure/firestore"
	"github.com/dappface/news-loader/interface/api/server/handler"
	"github.com/dappface/news-loader/interface/api/server/router"
	"github.com/dappface/news-loader/usecase"
)

type Registry interface {
	NewHandler() http.Handler
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func (s *registry) NewHandler() http.Handler {
	// repository
	postRepository := firestore.NewPost()
	latestTweetRepository := firestore.NewLatestTweet()

	// service
	rssService := service.NewRSS(postRepository)
	twitterService := service.NewTwitter(postRepository, latestTweetRepository)

	// usecase
	rssUseCase := usecase.NewRSS(rssService)
	twitterUseCase := usecase.NewTwitter(twitterService)

	// handler
	rssHandler := handler.NewRSS(rssUseCase)
	twitterHandler := handler.NewTwitter(twitterUseCase)

	// router
	r := router.NewRouter(rssHandler, twitterHandler)

	return r.CreateHandler()
}
