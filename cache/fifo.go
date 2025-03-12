package cache

import "fmt"

// FIFO defines a First-In-First-Out cache
type FIFO struct {
	capacity int
	cache    map[int]int
	order    []int
}

// NewFIFO creates a new FIFO cache with a specified capacity
func NewFIFO(capacity int) *FIFO {
	return &FIFO{
		capacity: capacity,
		cache:    make(map[int]int),
		order:    []int{},
	}
}

// Get retrieves the value for a key from the cache
func (f *FIFO) Get(key int) (int, bool) {
	value, exists := f.cache[key]
	return value, exists
}

// Put adds a new key-value pair into the cache, evicting the oldest if full
func (f *FIFO) Put(key int, value int) {
	if len(f.cache) == f.capacity {
		oldestKey := f.order[0]
		delete(f.cache, oldestKey)
		f.order = f.order[1:]
	}
	f.cache[key] = value
	f.order = append(f.order, key)
}

// Display prints the current state of the cache
func (f *FIFO) Display() {
	fmt.Println("Cache FIFO:", f.cache)
}

// Remove manually removes a key from the cache
func (f *FIFO) Remove(key int) {
	delete(f.cache, key)
	for i, k := range f.order {
		if k == key {
			f.order = append(f.order[:i], f.order[i+1:]...)
			break
		}
	}
}
