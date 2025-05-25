package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	shortenUrlHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "https://exm.pl/abc123")
	}

	longUrlHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "https://google.com/search?q=fluffy+cat")
	}

	metricHandler := func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "{\"youtube\" : 1}")
	}
	http.HandleFunc("POST /shorten", shortenUrlHandler)
	http.HandleFunc("GET /metric", metricHandler)
	http.HandleFunc("GET /{shortened_url}", longUrlHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
