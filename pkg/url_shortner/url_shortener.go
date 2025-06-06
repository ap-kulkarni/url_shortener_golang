package url_shortner

import (
	"fmt"
	"net/url"
)

const (
	ShortUrlLength = 7
	Host           = "http://localhost:8080/"
)

var urlsAggregate = InitShortenedUrlsAggregate()

func getFullUrl(shortUrlSegment string) string {
	return fmt.Sprintf("%s%s", Host, shortUrlSegment)
}

func ShortenUrl(urlToShorten *url.URL) string {
	var shortUrlSegment string
	shortUrlSegment = urlsAggregate.ShortenUrl(urlToShorten)
	return getFullUrl(shortUrlSegment)
}

func GetLongUrlFromShortUrl(shortUrlSegment string) string {
	return urlsAggregate.GetLongUrlFromShortUrl(shortUrlSegment)
}
