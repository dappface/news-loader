package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(r http.HandlerFunc, t http.HandlerFunc) Router {
	return &router{
		r,
		t,
	}
}

type Router interface {
	CreateHandler() http.Handler
}

type router struct {
	RssHandler     http.HandlerFunc
	TwitterHandler http.HandlerFunc
}

func (r *router) CreateHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/rss", r.RssHandler).Methods("POST")
	mux.HandleFunc("/twitter", r.TwitterHandler).Methods("POST")
	return mux
}
