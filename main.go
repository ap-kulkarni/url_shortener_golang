package main

import (
	"github.com/ap-kulkarni/url-shortener-assignment-infracloud/pkg/url_shortner"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("POST /shorten", url_shortner.ShortenUrlHandler)
	http.HandleFunc("GET /metric", url_shortner.MetricHandler)
	http.HandleFunc("GET /{shortened_url}", url_shortner.LongUrlHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
