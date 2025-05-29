package url_shortner

import (
	"fmt"
	"io"
	"net/http"
)

const (
	ShortUrlLength   = 7
	Host             = "http://localhost:8080"
	ShortUrlResponse = "{\"short_url\": \"%s/%s\"}"
)

var UrlsMap = map[string]string{}

func ShortenUrlHandler(w http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("error reading request body"))
		return
	}
	urlToShorten, err := GetUrlFromRequestBody(reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad request or no url found in the request"))
		return
	}
	if !ValidateUrl(urlToShorten) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("url not valid"))
		return
	}
	var shortUrlSegment string
	for {
		shortUrlSegment = GetRandomString(ShortUrlLength)
		if UrlsMap[shortUrlSegment] == "" {
			break
		}
	}
	UrlsMap[shortUrlSegment] = urlToShorten
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(fmt.Sprintf(ShortUrlResponse, Host, shortUrlSegment)))
}

func LongUrlHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "https://google.com/search?q=fluffy+cat")
}

func MetricHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "{\"youtube\" : 1}")
}
