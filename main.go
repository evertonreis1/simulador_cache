package main

import (
	"fmt"
	"simulador_cache/cache"

	"github.com/fatih/color"
)

func main() {

	// Cores
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	// Cabeçalho
	fmt.Println(cyan("====================================="))
	fmt.Println(cyan("          TESTING CACHE SYSTEM      "))
	fmt.Println(cyan("====================================="))

	// Testando o Cache FIFO de nível 1 com métricas
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

	// Testando o Cache LRU de nível 1 com métricas
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

	// Testando o Cache de múltiplos níveis (L1 e L2) com métricas
	fmt.Println("\n", yellow("Testing Multi-Level Cache with Metrics:"))
	multiCache := cache.NewMultiLevelCache(3, 3)

	multiCache.Put(1, 10)
	multiCache.Put(2, 20)
	multiCache.Put(3, 30)
	multiCache.Display()

	multiCache.Get(2)
	multiCache.Put(4, 40)
	multiCache.Display()

	// Mensagem final
	fmt.Println("\n", green("====================================="))
	fmt.Println(green("Cache system tests completed!"))
	fmt.Println(green("====================================="))
}
