package url_shortner

import (
	"container/heap"
	"testing"
)

func initDomainCounts() *DomainCounts {
	var domainCounts DomainCounts
	heap.Init(&domainCounts)
	return &domainCounts
}
func TestDomainCounts(t *testing.T) {
	t.Run("initial domain count length is 0", func(t *testing.T) {
		domainCounts := initDomainCounts()
		got := domainCounts.Len()
		want := 0
		assertCorrectMessage(t, got, want)
	})

	t.Run("push increases length by 1", func(t *testing.T) {
		domainCounts := initDomainCounts()
		heap.Push(domainCounts, &DomainCount{
			Domain: "google.com",
			Count:  1,
		})
		got := domainCounts.Len()
		want := 1
		assertCorrectMessage(t, got, want)
	})

	t.Run(" pop decreases length by 1", func(t *testing.T) {
		domainCounts := initDomainCounts()
		heap.Push(domainCounts, &DomainCount{
			Domain: "google.com",
			Count:  1,
		})
		_ = heap.Pop(domainCounts)
		got := domainCounts.Len()
		want := 0
		assertCorrectMessage(t, got, want)
	})
	t.Run("top element is with highest domain count", func(t *testing.T) {
		var domainCounts DomainCounts
		heap.Init(&domainCounts)
		domainCount1 := &DomainCount{
			Domain: "google.com",
			Count:  2,
		}
		domainCount2 := &DomainCount{
			Domain: "wikipedia.com",
			Count:  4,
		}
		domainCount3 := &DomainCount{
			Domain: "example.com",
			Count:  3,
		}
		heap.Push(&domainCounts, domainCount1)
		heap.Push(&domainCounts, domainCount2)
		heap.Push(&domainCounts, domainCount3)
		got := heap.Pop(&domainCounts).(*DomainCount).Domain
		want := domainCount2.Domain
		assertCorrectMessage(t, got, want)
	})
}
