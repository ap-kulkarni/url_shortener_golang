package url_shortner

import (
	"strings"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	t.Run("random string length is equal to length passed", func(t *testing.T) {
		got := len(GetRandomString(6))
		want := 6
		if got != want {
			t.Errorf("GetRandomString() = %d, want %d", got, want)
		}
	})

	t.Run("random string has characters only from predefined char set", func(t *testing.T) {
		randString := GetRandomString(10)
		for _, char := range randString {
			if !strings.Contains(randStringChars, string(char)) {
				t.Errorf("Random string contains characters outside of predefined character set")
			}
		}
	})
}
