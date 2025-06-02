package url_shortner

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"net/url"
	"sync"
)

const (
	ShortUrlLength = 7
	Host           = "http://localhost:8080/"
)

var urlsAggregate = &ShortenedUrlsAggregate{
	urlsMap:              make(map[string]string),
	lock:                 sync.Mutex{},
	shortenedUrlsHistory: bloom.NewWithEstimates(4_000_000_000_000, 0.1),
}

func getFullUrl(shortUrlSegment string) string {
	return fmt.Sprintf("%s%s", Host, shortUrlSegment)
}

func ShortenUrl(urlToShorten *url.URL) string {
	var shortUrlSegment string
	urlString := urlToShorten.String()
	shortUrlSegment = urlsAggregate.ShortenUrl(urlString)
	return getFullUrl(shortUrlSegment)
}

func GetLongUrlFromShortUrl(shortUrlSegment string) string {
	return urlsAggregate.GetLongUrlFromShortUrl(shortUrlSegment)
}
