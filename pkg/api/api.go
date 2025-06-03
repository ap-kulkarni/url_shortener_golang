package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ap-kulkarni/url-shortener-assignment-infracloud/pkg/url_shortner"
)

const shortenUrlResponse = "{\"short_url\": \"%s\"}"

func ShortenUrlHandler(w http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "error reading request body")
		return
	}
	urlToShorten, err := GetUrlFromRequestBody(reqBody)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if urlToShorten.Host == "" || urlToShorten.Scheme == "" {
		WriteErrorResponse(w, http.StatusBadRequest, "no valid url in request body")
	}
	shortUrl := url_shortner.ShortenUrl(urlToShorten)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(shortenUrlResponse, shortUrl)))
}

func GetLongUrlHandler(w http.ResponseWriter, req *http.Request) {
	shortUrlSegment := req.PathValue("shortened_url")
	longUrl := url_shortner.GetLongUrlFromShortUrl(shortUrlSegment)
	if longUrl == "" {
		WriteErrorResponse(w, http.StatusNotFound, "shortened url not found")
		return
	}
	w.Header().Set("Location", longUrl)
	w.WriteHeader(http.StatusMovedPermanently)
}

func MetricHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "{\"youtube\" : 1}")
}
