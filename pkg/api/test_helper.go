package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func assertCorrectStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status code %d, want %d", got, want)
	}
}

func assertCorrectMessage(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func readShortUrlFromResponse(t *testing.T, response []byte) string {
	t.Helper()
	var result map[string]any
	err := json.Unmarshal(response, &result)
	if err != nil {
		t.Errorf("Error unmarshalling body: %v", err)
	}
	return result["short_url"].(string)
}

func getShortenUrlRequestBody(urlToShorten string) string {
	return fmt.Sprintf(`{"url":"%s"}`, urlToShorten)
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

func executeMetricRequest() *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/metric", nil)
	w := httptest.NewRecorder()

	MetricHandler(w, req)
	return w
}

func extractShortUrlSegment(shortenedUrl string) string {
	parsedUrl, _ := url.Parse(shortenedUrl)
	return parsedUrl.Path
}
