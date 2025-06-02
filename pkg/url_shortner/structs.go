package url_shortner

import (
	"github.com/bits-and-blooms/bloom/v3"
	"sync"
)

type ShortenedUrlsAggregate struct {
	lock                 sync.Mutex
	urlsMap              map[string]string
	shortenedUrlsHistory *bloom.BloomFilter
}

func (s *ShortenedUrlsAggregate) GetLongUrlFromShortUrl(shortUrl string) string {
	return s.urlsMap[shortUrl]
}
func (s *ShortenedUrlsAggregate) ContainsLongUrl(urlToCheck string) bool {
	if !s.shortenedUrlsHistory.Test([]byte(urlToCheck)) {
		return false
	}
	for _, longUrl := range s.urlsMap {
		if longUrl == urlToCheck {
			return true
		}
	}
	return false
}

func (s *ShortenedUrlsAggregate) ContainsShortUrl(shortUrl string) bool {
	return s.GetLongUrlFromShortUrl(shortUrl) != ""
}

func (s *ShortenedUrlsAggregate) ShortenUrl(urlToShorten string) string {
	if s.shortenedUrlsHistory.Test([]byte(urlToShorten)) {
		for shortUrlSegment, longUrl := range s.urlsMap {
			if longUrl == urlToShorten {
				return shortUrlSegment
			}
		}
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	var shortUrlSegment string
	for {
		shortUrlSegment = GetRandomString(ShortUrlLength)
		if !s.ContainsShortUrl(shortUrlSegment) {
			break
		}
	}
	s.urlsMap[shortUrlSegment] = urlToShorten
	s.shortenedUrlsHistory.Add([]byte(urlToShorten))
	return shortUrlSegment
}
