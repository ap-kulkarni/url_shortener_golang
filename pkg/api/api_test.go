package api

import (
	"io"
	"net/http"
	"testing"
)

func TestShortenUrlHandler(t *testing.T) {
	t.Run("passing valid url gives short url", func(t *testing.T) {
		shortenReqBody := getShortenUrlRequestBody("https://www.google.com/search?q=fluffy+cat")
		w := executeShortenUrlRequest(shortenReqBody)
		assertCorrectStatusCode(t, w.Code, http.StatusOK)
		body, err := io.ReadAll(w.Result().Body)
		if err != nil {
			t.Errorf("Error reading body: %v", err)
		}
		shortUrl := readShortUrlFromResponse(t, body)
		if shortUrl == "" {
			t.Errorf("No short url in response")
		}
	})

	t.Run("passing invalid url gives err", func(t *testing.T) {
		shortenReqBody := getShortenUrlRequestBody("google")
		w := executeShortenUrlRequest(shortenReqBody)
		assertCorrectStatusCode(t, w.Code, http.StatusBadRequest)
	})

	t.Run("passing invalid request body gives err", func(t *testing.T) {
		shortenReqBody := `{"url_": "google"}`
		w := executeShortenUrlRequest(shortenReqBody)
		assertCorrectStatusCode(t, w.Code, http.StatusBadRequest)
	})
}

func TestGetLongUrlHandler(t *testing.T) {
	t.Run("passing valid short url redirects to long url", func(t *testing.T) {
		urlToShorten := "https://www.google.com/search?q=fluffy+cat"
		shortenReqBody := getShortenUrlRequestBody(urlToShorten)
		w := executeShortenUrlRequest(shortenReqBody)
		body, _ := io.ReadAll(w.Result().Body)
		shortUrl := readShortUrlFromResponse(t, body)
		shortUrlSegment := extractShortUrlSegment(shortUrl)
		w = executeGetLongUrlRequest(shortUrlSegment)
		assertCorrectStatusCode(t, w.Code, http.StatusMovedPermanently)
		if w.Header().Get("Location") != urlToShorten {
			t.Errorf("Received wrong location header: %v, expected: %v", w.Header(), urlToShorten)
		}
	})

	t.Run("passing invalid short url gives not found", func(t *testing.T) {
		shortUrlSegment := "/abc1234"
		w := executeGetLongUrlRequest(shortUrlSegment)
		assertCorrectStatusCode(t, w.Code, http.StatusNotFound)
	})
}
