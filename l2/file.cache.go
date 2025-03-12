package l2

import (
	"encoding/gob"
	"fmt"
	"os"
)

// FileCache defines the L2 cache that stores items in a file
type FileCache struct {
	capacity int
	filePath string
	cache    map[int]int
}

// NewFileCache creates a new FileCache with the given capacity and file path
func NewFileCache(capacity int) *FileCache {
	return &FileCache{
		capacity: capacity,
		filePath: "cache_l2.dat", // You can specify the file name here
		cache:    make(map[int]int),
	}
}

// Get retrieves a value from the L2 cache (disk)
func (f *FileCache) Get(key int) (int, bool) {
	f.loadCacheFromFile()
	value, exists := f.cache[key]
	return value, exists
}

// Put adds a value to the L2 cache and saves it to a file
func (f *FileCache) Put(key int, value int) {
	if len(f.cache) >= f.capacity {
		// Eviction logic can be implemented here
	}
	f.cache[key] = value
	f.saveCacheToFile()
}

// Remove deletes a key from the L2 cache and updates the file
func (f *FileCache) Remove(key int) {
	delete(f.cache, key)
	f.saveCacheToFile()
}

// Display shows the current state of the L2 cache
func (f *FileCache) Display() {
	fmt.Println(f.cache)
}

// saveCacheToFile saves the current cache to a file
func (f *FileCache) saveCacheToFile() {
	file, err := os.Create(f.filePath)
	if err != nil {
		fmt.Println("Error saving cache to file:", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(f.cache)
	if err != nil {
		fmt.Println("Error encoding cache to file:", err)
	}
}

// loadCacheFromFile loads the cache from the file
func (f *FileCache) loadCacheFromFile() {
	file, err := os.Open(f.filePath)
	if err != nil {
		fmt.Println("Error loading cache from file:", err)
		return
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&f.cache)
	if err != nil {
		fmt.Println("Error decoding cache from file:", err)
	}
}
