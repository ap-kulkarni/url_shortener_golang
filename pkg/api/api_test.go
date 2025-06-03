package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func readShortUrlFromResponse(t *testing.T, response []byte) string {
	t.Helper()
	var result map[string]any
	err := json.Unmarshal(response, &result)
	if err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	return result["short_url"].(string)
}

func executeShortenUrlRequest(shortenRequestBody string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(shortenRequestBody))
	w := httptest.NewRecorder()

	ShortenUrlHandler(w, req)
	return w
}

func executeGetLongUrlRequest(shortUrlSegment string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, shortUrlSegment, nil)
	req.SetPathValue("shortened_url", strings.TrimLeft(shortUrlSegment, "/"))
	w := httptest.NewRecorder()
	GetLongUrlHandler(w, req)
	return w
}

func extractShortUrlSegment(shortenedUrl string) string {
	parsedUrl, _ := url.Parse(shortenedUrl)
	return parsedUrl.Path
}

func TestShortenUrlHandler(t *testing.T) {
	t.Run("passing valid url gives short url", func(t *testing.T) {
		shortenReqBody := "{\"url\": \"https://www.google.com/search?q=fluffy+cat\"}"
		w := executeShortenUrlRequest(shortenReqBody)
		if w.Code != http.StatusOK {
			t.Errorf("Received wrong status code: %v, expected: %v", w.Code, http.StatusOK)
		}
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
		shortenReqBody := "{\"url\": \"google\"}"
		w := executeShortenUrlRequest(shortenReqBody)
		if w.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("shorten url api didn't return bad request for invalid url")
		}
	})

	t.Run("passing invalid request body gives err", func(t *testing.T) {
		shortenReqBody := "{\"url_\": \"google\"}"
		w := executeShortenUrlRequest(shortenReqBody)

		if w.Result().StatusCode != http.StatusBadRequest {
			t.Errorf("shorten url api didn't return bad request for invalid request body")
		}
	})
}

func TestGetLongUrlHandler(t *testing.T) {
	t.Run("passing valid short url redirects to long url", func(t *testing.T) {
		urlToShorten := "https://www.google.com/search?q=fluffy+cat"
		shortenReqBody := fmt.Sprintf("{\"url\": \"%s\"}", urlToShorten)
		w := executeShortenUrlRequest(shortenReqBody)
		body, _ := io.ReadAll(w.Result().Body)
		shortUrl := readShortUrlFromResponse(t, body)
		shortUrlSegment := extractShortUrlSegment(shortUrl)
		w = executeGetLongUrlRequest(shortUrlSegment)
		if w.Code != http.StatusMovedPermanently {
			t.Errorf("Received wrong status code: %v, expected: %v", w.Code, http.StatusMovedPermanently)
		}
		if w.Header().Get("Location") != urlToShorten {
			t.Errorf("Received wrong location header: %v, expected: %v", w.Header(), urlToShorten)
		}
	})

	t.Run("passing invalid short url gives not found", func(t *testing.T) {
		shortUrlSegment := "/abc1234"
		w := executeGetLongUrlRequest(shortUrlSegment)
		if w.Code != http.StatusNotFound {
			t.Errorf("Received wrong status code: %v, expected: %v", w.Code, http.StatusNotFound)
		}
	})
}
