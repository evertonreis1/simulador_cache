package cache

type Cache interface {
	Get(key int) (int, bool)
	Put(key int, value int)
	Display()
	Remove(key int)
	ShowMetrics()
}
