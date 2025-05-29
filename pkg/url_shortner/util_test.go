package url_shortner

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetRandomString(t *testing.T) {
	t.Run("passing zero length return blank string", func(t *testing.T) {
		got := GetRandomString(0)
		want := ""
		assertCorrectMessage(t, got, want)
	})
	t.Run("random string length is equal to length passed", func(t *testing.T) {
		got := len(GetRandomString(6))
		want := 6
		assertCorrectMessage(t, got, want)
	})
}

func TestGetUrlFromRequestBody(t *testing.T) {
	testUrl := "http://example.com"
	t.Run("passing valid url", func(t *testing.T) {
		validRequestBody := []byte(fmt.Sprintf("{\"url\": \"%s\"}", testUrl))
		got, _ := GetUrlFromRequestBody(validRequestBody)
		want := testUrl
		assertCorrectMessage(t, got, want)
	})
	t.Run("passing invalid url", func(t *testing.T) {
		invalidRequestBody := []byte(fmt.Sprintf("{\"dummy\": \"%s\"}", testUrl))
		_, err := GetUrlFromRequestBody(invalidRequestBody)
		if err == nil {
			t.Errorf("expected error when invalid url is passed")
		}
	})
}
