package url_shortner

import (
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
