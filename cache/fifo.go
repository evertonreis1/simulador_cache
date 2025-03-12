package cache

import (
	"fmt"
)

type FIFO struct {
	capacity int
	cache    map[int]int
	order    []int
	metrics  CacheMetrics
}

func NewFIFO(capacity int) *FIFO {
	return &FIFO{
		capacity: capacity,
		cache:    make(map[int]int),
		order:    []int{},
		metrics:  CacheMetrics{},
	}
}

func (f *FIFO) Get(key int) (int, bool) {
	f.metrics.TotalGets++
	value, exists := f.cache[key]
	if exists {
		f.metrics.Hits++
	} else {
		f.metrics.Misses++
	}
	return value, exists
}

func (f *FIFO) Put(key int, value int) {
	if len(f.cache) == f.capacity {

		oldestKey := f.order[0]
		delete(f.cache, oldestKey)
		f.order = f.order[1:]
		f.metrics.Removals++
	}
	f.cache[key] = value
	f.order = append(f.order, key)
}

func (f *FIFO) Remove(key int) {
	delete(f.cache, key)

	for i, k := range f.order {
		if k == key {
			f.order = append(f.order[:i], f.order[i+1:]...)
			break
		}
	}
}

func (f *FIFO) Display() {
	fmt.Println("Cache FIFO:", f.cache)
}

func (f *FIFO) ShowMetrics() {
	f.metrics.ShowMetrics()
}
