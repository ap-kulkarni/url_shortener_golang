package api

import (
	"fmt"
	"testing"
)

func TestGetUrlFromRequestBody(t *testing.T) {
	testUrl := "http://test.com"
	t.Run("passing valid url", func(t *testing.T) {
		validRequestBody := []byte(fmt.Sprintf("{\"url\": \"%s\"}", testUrl))
		got, _ := GetUrlFromRequestBody(validRequestBody)
		want := testUrl
		assertCorrectMessage(t, got.String(), want)
	})
	t.Run("passing invalid url", func(t *testing.T) {
		invalidRequestBody := []byte(fmt.Sprintf("{\"dummy\": \"%s\"}", testUrl))
		_, err := GetUrlFromRequestBody(invalidRequestBody)
		if err == nil {
			t.Errorf("expected error when invalid url is passed")
		}
	})
}
