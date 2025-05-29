package url_shortner

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"net/url"
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

func GetUrlFromRequestBody(requestBody []byte) (string, error) {
	var parsed map[string]any
	err := json.Unmarshal(requestBody, &parsed)
	if err != nil {
		return "", err
	}
	parsedUrl := parsed["url"]
	if parsedUrl == nil {
		return "", errors.New("no url in request body")
	}
	return parsedUrl.(string), nil
}

func ValidateUrl(urlToValidate string) bool {
	if urlToValidate == "" {
		return false
	}
	_, err := url.Parse(urlToValidate)
	return err == nil
}
