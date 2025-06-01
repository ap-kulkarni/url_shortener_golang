package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const errorResponse = "{\"error\": \"%s\"}"

func WriteErrorResponse(w http.ResponseWriter, status int, error string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(fmt.Sprintf(errorResponse, error)))
}

func GetUrlFromRequestBody(requestBody []byte) (*url.URL, error) {
	var parsedBody map[string]any
	err := json.Unmarshal(requestBody, &parsedBody)
	if err != nil {
		return nil, err
	}
	extractedUrl := parsedBody["url"]
	if extractedUrl == nil {
		return nil, errors.New("no url in request body")
	}
	return url.Parse(extractedUrl.(string))
}
