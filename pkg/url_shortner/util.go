package url_shortner

import (
	"math/rand/v2"
	"strings"
)

const randStringChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRandomString(length int) string {
	if length <= 0 {
		return ""
	}
	randomString := strings.Builder{}
	for range length {
		index := rand.N(len(randStringChars))
		randomString.WriteByte(randStringChars[index])
	}
	return randomString.String()
}
