package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simulador_cache/cache"
	"time"

	"github.com/fatih/color"
)

var metricsL1 = cache.NewMetrics()
var metricsL2 = cache.NewMetrics()

func getMetrics(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"l1": map[string]int{
			"hits":   metricsL1.HitCount,
			"misses": metricsL1.MissCount,
		},
		"l2": map[string]int{
			"hits":   metricsL2.HitCount,
			"misses": metricsL2.MissCount,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {

	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println(cyan("====================================="))
	fmt.Println(cyan("          TESTING CACHE SYSTEM      "))
	fmt.Println(cyan("====================================="))

	fmt.Println("\n", yellow("Testing FIFO Cache with Metrics:"))
	fifoCache := cache.NewFIFO(3)

	fifoCache.Put(1, 10)
	fifoCache.Put(2, 20)
	fifoCache.Put(3, 30)
	fifoCache.Display()

	fifoCache.Get(2)
	fifoCache.Put(4, 40)
	fifoCache.Display()
	fifoCache.ShowMetrics()

	fmt.Println("\n", yellow("Testing LRU Cache with Metrics:"))
	lruCache := cache.NewLRU(3)

	lruCache.Put(1, 10)
	lruCache.Put(2, 20)
	lruCache.Put(3, 30)
	lruCache.Display()

	lruCache.Get(1)
	lruCache.Put(4, 40)
	lruCache.Display()
	lruCache.ShowMetrics()

	fmt.Println("\n", yellow("Testing LFU Cache with Metrics:"))
	lfuCache := cache.NewLFU(3)

	lfuCache.Put(1, 10)
	lfuCache.Put(2, 20)
	lfuCache.Put(3, 30)
	lfuCache.Display()

	lfuCache.Get(1)
	lfuCache.Put(4, 40)
	lfuCache.Display()
	lfuCache.ShowMetrics()

	fmt.Println("\n", yellow("Testing Multi-Level Cache with Metrics:"))
	multiCache := cache.NewMultiLevelCache(3, 3)

	multiCache.Put(1, 10)
	multiCache.Put(2, 20)
	multiCache.Put(3, 30)
	multiCache.Display()

	multiCache.Get(2)
	multiCache.Put(4, 40)
	multiCache.Display()

	// ===========================
	// Testando Cache TTL
	// ===========================
	fmt.Println("\n", yellow("Testing TTL Cache with Expiration:"))
	ttlCache := cache.NewTTLCache(3)

	ttlCache.Put(1, 10, 3*time.Second)
	ttlCache.Put(2, 20, 5*time.Second)
	ttlCache.Put(3, 30, 1*time.Second)

	fmt.Println("\nInitial TTL Cache State:")
	ttlCache.Display()

	time.Sleep(2 * time.Second)
	fmt.Println("\nAfter 2 seconds:")
	ttlCache.Display()

	time.Sleep(3 * time.Second)
	fmt.Println("\nAfter 5 seconds (some items should be expired):")
	ttlCache.Display()

	fmt.Println("\n", green("====================================="))
	fmt.Println(green("Cache system tests completed!"))
	fmt.Println(green("====================================="))

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/metrics", getMetrics)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
