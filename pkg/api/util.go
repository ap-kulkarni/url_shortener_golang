package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ap-kulkarni/url_shortener_golang/pkg/url_shortner"
	"net/http"
	"net/url"
	"strings"
)

const errorResponse = `{"error": "%s"}`

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

func GetTopDomainDataInJson(domainCounts []url_shortner.DomainCount) string {
	domainDataJson := strings.Builder{}
	domainDataJson.WriteString("{")
	for i, domain := range domainCounts {
		if i < len(domainCounts)-1 {
			domainDataJson.WriteString(fmt.Sprintf(`"%s": %d,`, domain.Domain, domain.Count))
		}
	}
	lastDomain := domainCounts[len(domainCounts)-1]
	domainDataJson.WriteString(fmt.Sprintf(`"%s": %d}`, lastDomain.Domain, lastDomain.Count))
	return domainDataJson.String()
}
