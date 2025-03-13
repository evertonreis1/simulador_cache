package cache

import "fmt"

type LFU struct {
	capacity  int
	cache     map[int]int
	frequency map[int]int
	order     map[int]int
	metrics   CacheMetrics
}

func NewLFU(capacity int) *LFU {
	return &LFU{
		capacity:  capacity,
		cache:     make(map[int]int),
		frequency: make(map[int]int),
		order:     make(map[int]int),
		metrics:   CacheMetrics{},
	}
}

func (l *LFU) Get(key int) (int, bool) {
	l.metrics.TotalGets++
	value, exists := l.cache[key]
	if exists {
		l.metrics.Hits++
		l.frequency[key]++
	} else {
		l.metrics.Misses++
	}
	return value, exists
}

func (l *LFU) Put(key int, value int) {
	if len(l.cache) == l.capacity {
		l.removeLeastFrequentlyUsed()
	}
	l.cache[key] = value
	l.frequency[key] = 1
	l.order[key] = l.metrics.TotalGets
}

func (l *LFU) removeLeastFrequentlyUsed() {
	leastFrequent := -1
	leastFrequency := int(^uint(0) >> 1) // Max int value

	for key, freq := range l.frequency {
		if freq < leastFrequency {
			leastFrequency = freq
			leastFrequent = key
		}
	}

	delete(l.cache, leastFrequent)
	delete(l.frequency, leastFrequent)
	delete(l.order, leastFrequent)
}

func (l *LFU) Remove(key int) {
	delete(l.cache, key)
	delete(l.frequency, key)
	delete(l.order, key)
}

func (l *LFU) Display() {
	fmt.Println("Cache LFU:", l.cache)
}

func (l *LFU) ShowMetrics() {
	l.metrics.ShowMetrics()
}
