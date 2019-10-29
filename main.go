package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dappface/news-loader/registry"
	_ "github.com/GoogleCloudPlatform/berglas/pkg/auto"
)

var h http.Handler
var port string

func init() {
	r := registry.NewRegistry()
	h = r.NewHandler()

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}

func main() {
	http.Handle("/", h)

	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
