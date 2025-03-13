package cache

import (
	"fmt"
)

type CacheMetrics struct {
	Hits        int
	Misses      int
	Removals    int
	TotalGets   int
	HitRate     float64
	MissRate    float64
	RemovalRate float64
}

func (metrics *CacheMetrics) ShowMetrics() {
	metrics.HitRate = float64(metrics.Hits) / float64(metrics.TotalGets) * 100
	metrics.MissRate = float64(metrics.Misses) / float64(metrics.TotalGets) * 100
	metrics.RemovalRate = float64(metrics.Removals) / float64(metrics.TotalGets) * 100

	fmt.Println("Metrics:")
	fmt.Printf("Hit Rate: %.2f%%\n", metrics.HitRate)
	fmt.Printf("Miss Rate: %.2f%%\n", metrics.MissRate)
	fmt.Printf("Removal Rate: %.2f%%\n", metrics.RemovalRate)
}
