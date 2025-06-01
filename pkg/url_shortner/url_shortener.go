package url_shortner

import (
	"fmt"
	"net/url"
)

const (
	ShortUrlLength = 7
	Host           = "http://localhost:8080/"
)

var UrlsMap = map[string]string{}

func getFullUrl(shortUrlSegment string) string {
	return fmt.Sprintf("%s%s", Host, shortUrlSegment)
}

func getUniqueShortSegment() string {
	var shortUrlSegment string
	for {
		shortUrlSegment = GetRandomString(ShortUrlLength)
		if UrlsMap[shortUrlSegment] == "" {
			break
		}
	}
	return shortUrlSegment
}

func ShortenUrl(urlToShorten *url.URL) string {
	var shortUrlSegment string
	urlString := urlToShorten.String()

	for key, value := range UrlsMap {
		if value == urlString {
			shortUrlSegment = key
		}
	}
	if shortUrlSegment == "" {
		shortUrlSegment = getUniqueShortSegment()
		UrlsMap[shortUrlSegment] = urlToShorten.String()
	}

	return getFullUrl(shortUrlSegment)
}

func GetLongUrlFromShortUrl(shortUrlSegment string) string {
	return UrlsMap[shortUrlSegment]
}
