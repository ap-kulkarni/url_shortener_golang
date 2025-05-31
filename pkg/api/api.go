package api

import (
	"github.com/ap-kulkarni/url-shortener-assignment-infracloud/pkg/url_shortner"
	"io"
	"net/http"
)

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
	shortUrl := url_shortner.ShortenUrl(urlToShorten)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(shortUrl))
}

func LongUrlHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "https://google.com/search?q=fluffy+cat")
}

func MetricHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "{\"youtube\" : 1}")
}
