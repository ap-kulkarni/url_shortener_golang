package url_shortner

import (
	"container/heap"
	"github.com/bits-and-blooms/bloom/v3"
	"net/url"
	"sync"
)

type DomainCount struct {
	Domain string
	Count  int
	index  int
}

/**************************************************************************
* DomainCounts is designed as implementation of heap interface. So that,  *
* first n will always give top n domains converted.                       *
* Implementation referred from the PriorityQueue example from             *
* https://pkg.go.dev/container/heap#example-package-PriorityQueue         *
***************************************************************************/

type DomainCounts []*DomainCount

func (d *DomainCounts) Len() int {
	return len(*d)
}

func (d *DomainCounts) Less(i, j int) bool {
	return (*d)[i].Count > (*d)[j].Count
}

func (d *DomainCounts) Swap(i, j int) {
	(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	(*d)[i].index = i
	(*d)[j].index = j
}

func (d *DomainCounts) Push(x any) {
	n := len(*d)
	domainCount := x.(*DomainCount)
	domainCount.index = n
	*d = append(*d, domainCount)
}

func (d *DomainCounts) Pop() any {
	old := *d
	n := len(old)
	domainCount := old[n-1]
	old[n-1] = nil
	*d = old[0 : n-1]
	domainCount.index = -1
	return domainCount
}

func (d *DomainCounts) Update(x *DomainCount, domain string, count int) {
	x.Domain = domain
	x.Count = count
	heap.Fix(d, x.index)
}

type ShortenedUrlsAggregate struct {
	lock                  sync.Mutex
	urlsMap               map[string]string
	shortenedUrlsHistory  *bloom.BloomFilter
	shortenedDomainsStats *DomainCounts
}

func InitShortenedUrlsAggregate() *ShortenedUrlsAggregate {
	urlsAggregate := &ShortenedUrlsAggregate{
		urlsMap:               make(map[string]string),
		lock:                  sync.Mutex{},
		shortenedUrlsHistory:  bloom.NewWithEstimates(4_000_000_000_000, 0.1),
		shortenedDomainsStats: &DomainCounts{},
	}
	heap.Init(urlsAggregate.shortenedDomainsStats)
	return urlsAggregate
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

func (s *ShortenedUrlsAggregate) ShortenUrl(urlToShorten *url.URL) string {
	urlString := urlToShorten.String()
	if s.shortenedUrlsHistory.Test([]byte(urlString)) {
		for shortUrlSegment, longUrl := range s.urlsMap {
			if longUrl == urlString {
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
	s.urlsMap[shortUrlSegment] = urlString
	s.shortenedUrlsHistory.Add([]byte(urlString))
	s.updateDomainCount(urlToShorten.Host)
	return shortUrlSegment
}

func (s *ShortenedUrlsAggregate) updateDomainCount(domain string) {
	domainStats := s.shortenedDomainsStats
	for _, d := range *domainStats {
		if d.Domain == domain {
			domainStats.Update(d, domain, d.Count+1)
			return
		}
	}
	domainCount := &DomainCount{domain, 1, 0}
	heap.Push(domainStats, domainCount)
}
