package handler

import (
	"net/http"

	"github.com/dappface/news-loader/usecase"
)

func NewTwitter(twitterUseCase usecase.Twitter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		twitterUseCase.LoadAll()
		w.WriteHeader(200)
	}
}
