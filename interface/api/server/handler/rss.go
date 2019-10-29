package handler

import (
	"net/http"

	"github.com/dappface/news-loader/usecase"
)

func NewRSS(rssUseCase usecase.RSS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rssUseCase.LoadAll()
		w.WriteHeader(200)
	}
}
