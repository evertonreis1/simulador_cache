package cache

import (
	"fmt"
	"simulador_cache/l2"
)

// Metrics holds cache metrics
type Metrics struct {
	HitCount  int
	MissCount int
}

// NewMetrics creates a new Metrics instance
func NewMetrics() *Metrics {
	return &Metrics{}
}

// DisplayMetrics displays the metrics
func (m *Metrics) DisplayMetrics() {
	fmt.Printf("Hits: %d, Misses: %d\n", m.HitCount, m.MissCount)
}

// MultiLevelCache defines a multi-level cache with L1 (RAM) and L2 (Disk)
type MultiLevelCache struct {
	L1         Cache         // Cache de Nível 1 (RAM)
	L2         *l2.FileCache // Cache de Nível 2 (Disk)
	L1Capacity int
	L2Capacity int
	metricsL1  *Metrics // Métricas para L1
	metricsL2  *Metrics // Métricas para L2
}

// NewMultiLevelCache creates a new MultiLevelCache with given capacities for L1 and L2
func NewMultiLevelCache(L1Capacity, L2Capacity int) *MultiLevelCache {
	return &MultiLevelCache{
		L1:         NewFIFO(L1Capacity), // L1 como FIFO, você pode mudar para LRU
		L2:         l2.NewFileCache(L2Capacity),
		L1Capacity: L1Capacity,
		L2Capacity: L2Capacity,
		metricsL1:  NewMetrics(),
		metricsL2:  NewMetrics(),
	}
}

// Get tries to get a value from L1, if not found, it looks in L2
func (m *MultiLevelCache) Get(key int) (int, bool) {
	if value, exists := m.L1.Get(key); exists {
		m.metricsL1.HitCount++
		return value, true
	}

	if value, exists := m.L2.Get(key); exists {
		m.metricsL2.HitCount++
		return value, true
	}

	m.metricsL1.MissCount++
	return 0, false
}

// Put adds a value to L1 and moves items to L2 if L1 is full
func (m *MultiLevelCache) Put(key int, value int) {
	if len(m.L1.(*FIFO).cache) == m.L1Capacity {
		// Move item from L1 to L2 when L1 is full
		oldestKey := m.L1.(*FIFO).order[0]
		oldestValue := m.L1.(*FIFO).cache[oldestKey]
		m.L2.Put(oldestKey, oldestValue)
		m.L1.Remove(oldestKey) // Remove the item from L1
	}
	m.L1.Put(key, value)
}

// Remove deletes a key from L1 and also from L2 if it exists
func (m *MultiLevelCache) Remove(key int) {
	m.L1.Remove(key)
	m.L2.Remove(key)
}

// Display shows the current state of both L1 and L2 caches and their metrics
func (m *MultiLevelCache) Display() {
	fmt.Println("Cache L1 (FIFO):", m.L1.(*FIFO).cache)
	m.metricsL1.DisplayMetrics()

	fmt.Println("Cache L2 (Disk):")
	m.L2.Display()
	m.metricsL2.DisplayMetrics()
}
