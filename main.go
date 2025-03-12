package main

import (
	"fmt"
	"simulador_cache/cache"
)

func main() {

	fmt.Println("Testing FIFO Cache with Metrics:")

	fifoCache := cache.NewFIFO(3)

	fifoCache.Put(1, 10)
	fifoCache.Put(2, 20)
	fifoCache.Put(3, 30)
	fifoCache.Display()

	fifoCache.Get(2)
	fifoCache.Put(4, 40)
	fifoCache.Display()
	fifoCache.ShowMetrics()

	fmt.Println("\nTesting LRU Cache with Metrics:")

	lruCache := cache.NewLRU(3)

	lruCache.Put(1, 10)
	lruCache.Put(2, 20)
	lruCache.Put(3, 30)
	lruCache.Display()

	lruCache.Get(1)
	lruCache.Put(4, 40)
	lruCache.Display()
	lruCache.ShowMetrics()
}
