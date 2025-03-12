package main

import (
	"fmt"
	"simulador_cache/cache"
)

func main() {
	fmt.Println("Testing FIFO:")
	fifoCache := cache.NewFIFO(3)
	fifoCache.Put(1, 10)
	fifoCache.Put(2, 20)
	fifoCache.Put(3, 30)
	fifoCache.Display()

	fifoCache.Put(4, 40)
	fifoCache.Display()

	value, exists := fifoCache.Get(2)
	if exists {
		fmt.Println("Value for key 2:", value)
	} else {
		fmt.Println("Key 2 not found.")
	}

	fifoCache.Remove(2)
	fmt.Println("After removing key 2:")
	fifoCache.Display()

	fmt.Println("\nTesting LRU:")
	lruCache := cache.NewLRU(3)
	lruCache.Put(1, 10)
	lruCache.Put(2, 20)
	lruCache.Put(3, 30)
	lruCache.Display()

	lruCache.Get(1)
	lruCache.Put(4, 40)
	lruCache.Display()

	value, exists = lruCache.Get(2)
	if exists {
		fmt.Println("Value for key 2:", value)
	} else {
		fmt.Println("Key 2 not found.")
	}

	lruCache.Remove(2)
	fmt.Println("After removing key 2:")
	lruCache.Display()
}
