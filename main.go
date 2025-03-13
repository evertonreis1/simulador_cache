package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simulador_cache/cache"

	"github.com/fatih/color"
)

var metricsL1 = cache.NewMetrics()
var metricsL2 = cache.NewMetrics()

// Função para exibir as métricas em formato JSON via HTTP
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

	// Testando o Cache LFU de nível 1 com métricas
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

	// Inicia o servidor web para o dashboard
	http.Handle("/", http.FileServer(http.Dir("./web"))) // Serve arquivos estáticos da pasta web
	http.HandleFunc("/metrics", getMetrics)              // Endpoint para métricas de cache

	// Inicia o servidor na porta 8080
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
