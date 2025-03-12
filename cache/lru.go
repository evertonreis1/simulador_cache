package cache

import "fmt"

// LRU defines a Least Recently Used cache
type LRU struct {
	capacity int
	cache    map[int]int
	order    []int
}

// NewLRU creates a new LRU cache with a specified capacity
func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		cache:    make(map[int]int),
		order:    []int{},
	}
}

// Get retrieves the value for a key from the cache
func (l *LRU) Get(key int) (int, bool) {
	value, exists := l.cache[key]
	if exists {
		l.moveToMostRecent(key)
	}
	return value, exists
}

// Put adds a new key-value pair into the cache, evicting the least recently used if full
func (l *LRU) Put(key int, value int) {
	if len(l.cache) == l.capacity {
		l.removeLeastRecent()
	}
	l.cache[key] = value
	l.order = append(l.order, key)
}

// moveToMostRecent moves a key to the end of the order to mark it as recently used
func (l *LRU) moveToMostRecent(key int) {
	for i, k := range l.order {
		if k == key {
			l.order = append(l.order[:i], l.order[i+1:]...)
			break
		}
	}
	l.order = append(l.order, key)
}

// removeLeastRecent removes the least recently used key from the cache
func (l *LRU) removeLeastRecent() {
	leastRecent := l.order[0]
	delete(l.cache, leastRecent)
	l.order = l.order[1:]
}

// Display prints the current state of the cache
func (l *LRU) Display() {
	fmt.Println("Cache LRU:", l.cache)
}

// Remove manually removes a key from the cache
func (l *LRU) Remove(key int) {
	delete(l.cache, key)
	for i, k := range l.order {
		if k == key {
			l.order = append(l.order[:i], l.order[i+1:]...)
			break
		}
	}
}
