package cache

import (
	"fmt"
)

type LRU struct {
	capacity int
	cache    map[int]int
	order    []int
	metrics  CacheMetrics
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		cache:    make(map[int]int),
		order:    []int{},
		metrics:  CacheMetrics{},
	}
}

func (l *LRU) Get(key int) (int, bool) {
	l.metrics.TotalGets++
	value, exists := l.cache[key]
	if exists {
		l.metrics.Hits++
		l.moveToMostRecent(key)
	} else {
		l.metrics.Misses++
	}
	return value, exists
}

func (l *LRU) Put(key int, value int) {
	if len(l.cache) == l.capacity {

		l.removeLeastRecent()
	}
	l.cache[key] = value
	l.order = append(l.order, key)
}

func (l *LRU) moveToMostRecent(key int) {
	for i, k := range l.order {
		if k == key {
			l.order = append(l.order[:i], l.order[i+1:]...)
			break
		}
	}
	l.order = append(l.order, key)
}

func (l *LRU) removeLeastRecent() {
	leastRecent := l.order[0]
	delete(l.cache, leastRecent)
	l.order = l.order[1:]
}

func (l *LRU) Remove(key int) {
	delete(l.cache, key)

	for i, k := range l.order {
		if k == key {
			l.order = append(l.order[:i], l.order[i+1:]...)
			break
		}
	}
}

func (l *LRU) Display() {
	fmt.Println("Cache LRU:", l.cache)
}

func (l *LRU) ShowMetrics() {
	l.metrics.ShowMetrics()
}
