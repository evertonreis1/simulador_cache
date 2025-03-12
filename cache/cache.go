package cache

// Cache defines the basic interface for a cache system
type Cache interface {
	Get(key int) (int, bool)
	Put(key int, value int)
	Display()
	Remove(key int)
}
