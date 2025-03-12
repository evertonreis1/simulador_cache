package cache

import (
	"fmt"
	"simulador_cache/l2"
)

type Metrics struct {
	HitCount  int
	MissCount int
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) DisplayMetrics() {
	fmt.Printf("Hits: %d, Misses: %d\n", m.HitCount, m.MissCount)
}

type MultiLevelCache struct {
	L1         Cache         // Cache de Nível 1 (RAM)
	L2         *l2.FileCache // Cache de Nível 2 (Disk)
	L1Capacity int
	L2Capacity int
	metricsL1  *Metrics
	metricsL2  *Metrics
}

func NewMultiLevelCache(L1Capacity, L2Capacity int) *MultiLevelCache {
	return &MultiLevelCache{
		L1:         NewFIFO(L1Capacity),
		L2:         l2.NewFileCache(L2Capacity),
		L1Capacity: L1Capacity,
		L2Capacity: L2Capacity,
		metricsL1:  NewMetrics(),
		metricsL2:  NewMetrics(),
	}
}

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

func (m *MultiLevelCache) Put(key int, value int) {
	if len(m.L1.(*FIFO).cache) == m.L1Capacity {

		oldestKey := m.L1.(*FIFO).order[0]
		oldestValue := m.L1.(*FIFO).cache[oldestKey]
		m.L2.Put(oldestKey, oldestValue)
		m.L1.Remove(oldestKey)
	}
	m.L1.Put(key, value)
}

func (m *MultiLevelCache) Remove(key int) {
	m.L1.Remove(key)
	m.L2.Remove(key)
}

func (m *MultiLevelCache) Display() {
	fmt.Println("Cache L1 (FIFO):", m.L1.(*FIFO).cache)
	m.metricsL1.DisplayMetrics()

	fmt.Println("Cache L2 (Disk):")
	m.L2.Display()
	m.metricsL2.DisplayMetrics()
}
