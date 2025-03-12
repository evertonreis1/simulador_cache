package main

import (
    "fmt"
)

type Cache interface {
    Get(key int) (int, bool)
    Put(key int, value int)
    Display()
}

type FIFO struct {
    capacity int
    cache    map[int]int
    order    []int
}

func NewFIFO(capacity int) *FIFO {
    return &FIFO{
        capacity: capacity,
        cache:    make(map[int]int),
        order:    []int{},
    }
}

func (f *FIFO) Get(key int) (int, bool) {
    value, exists := f.cache[key]
    return value, exists
}

func (f *FIFO) Put(key int, value int) {
    if len(f.cache) == f.capacity {
        oldestKey := f.order[0]
        delete(f.cache, oldestKey)
        f.order = f.order[1:]
    }
    f.cache[key] = value
    f.order = append(f.order, key)
}

func (f *FIFO) Display() {
    fmt.Println("Cache FIFO:", f.cache)
}

type LRU struct {
    capacity int
    cache    map[int]int
    order    []int
}

func NewLRU(capacity int) *LRU {
    return &LRU{
        capacity: capacity,
        cache:    make(map[int]int),
        order:    []int{},
    }
}

func (l *LRU) Get(key int) (int, bool) {
    value, exists := l.cache[key]
    if exists {
        l.moveToMostRecent(key)
    }
    return value, exists
}

func (l *LRU) Put(key int, value int) {
    if len(l.cache) == l.capacity {
        l.removeLeastRecent()
    }
    l.cache[key] = value
    l.order = append(l.order, key)
}

func (l *LRU) moveToMostRecent(key int) {
    
    for i, k := range l.order {
        if k == key {
            l.order = append(l.order[:i], l.order[i+1:]...)
            break
        }
    }
    
    l.order = append(l.order, key)
}

func (l *LRU) removeLeastRecent() {
    leastRecent := l.order[0]
    delete(l.cache, leastRecent)
    l.order = l.order[1:]
}

func (l *LRU) Display() {
    fmt.Println("Cache LRU:", l.cache)
}

func main() {
    fmt.Println("Testando FIFO:")
    fifoCache := NewFIFO(3)
    fifoCache.Put(1, 10)
    fifoCache.Put(2, 20)
    fifoCache.Put(3, 30)
    fifoCache.Display()

    fifoCache.Put(4, 40)
    fifoCache.Display()

   
    value, exists := fifoCache.Get(2)
    if exists {
        fmt.Println("Valor de chave 2:", value)
    } else {
        fmt.Println("Chave 2 não encontrada.")
    }


    fmt.Println("\nTestando LRU:")
    lruCache := NewLRU(3)
    lruCache.Put(1, 10)
    lruCache.Put(2, 20)
    lruCache.Put(3, 30)
    lruCache.Display()

    lruCache.Get(1)
    lruCache.Put(4, 40)
    lruCache.Display()

 
    value, exists = lruCache.Get(2)
    if exists {
        fmt.Println("Valor de chave 2:", value)
    } else {
        fmt.Println("Chave 2 não encontrada.")
    }
}
