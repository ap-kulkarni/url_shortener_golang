package main

import (
	"github.com/ap-kulkarni/url-shortener-assignment-infracloud/pkg/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("POST /shorten", api.ShortenUrlHandler)
	http.HandleFunc("GET /metric", api.MetricHandler)
	http.HandleFunc("GET /{shortened_url}", api.GetLongUrlHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
