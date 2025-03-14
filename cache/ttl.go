package cache

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	Value      int
	Expiration int64
}

type TTLCache struct {
	capacity int
	cache    map[int]CacheItem
	mutex    sync.Mutex
}

func NewTTLCache(capacity int) *TTLCache {
	ttlCache := &TTLCache{
		capacity: capacity,
		cache:    make(map[int]CacheItem),
	}
	go ttlCache.cleanupExpiredItems()
	return ttlCache
}

func (c *TTLCache) Put(key int, value int, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.cache) >= c.capacity {
		c.evict()
	}

	c.cache[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
}

func (c *TTLCache) Get(key int) (int, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, found := c.cache[key]
	if !found || time.Now().UnixNano() > item.Expiration {
		delete(c.cache, key)
		return 0, false
	}

	return item.Value, true
}

func (c *TTLCache) evict() {
	var oldestKey int64 = time.Now().UnixNano()
	var keyToRemove int

	for key, item := range c.cache {
		if item.Expiration < oldestKey {
			oldestKey = item.Expiration
			keyToRemove = key
		}
	}
	delete(c.cache, keyToRemove)
}

func (c *TTLCache) cleanupExpiredItems() {
	for {
		time.Sleep(time.Second)
		c.mutex.Lock()
		for key, item := range c.cache {
			if time.Now().UnixNano() > item.Expiration {
				delete(c.cache, key)
			}
		}
		c.mutex.Unlock()
	}
}

func (c *TTLCache) Display() {
	fmt.Println("Cache TTL:", c.cache)
}
