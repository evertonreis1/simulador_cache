package l2

import (
	"encoding/gob"
	"fmt"
	"os"
)

type FileCache struct {
	capacity int
	filePath string
	cache    map[int]int
}

func NewFileCache(capacity int) *FileCache {
	return &FileCache{
		capacity: capacity,
		filePath: "cache_l2.dat",
		cache:    make(map[int]int),
	}
}

func (f *FileCache) Get(key int) (int, bool) {
	f.loadCacheFromFile()
	value, exists := f.cache[key]
	return value, exists
}

func (f *FileCache) Put(key int, value int) {
	if len(f.cache) >= f.capacity {

	}
	f.cache[key] = value
	f.saveCacheToFile()
}

func (f *FileCache) Remove(key int) {
	delete(f.cache, key)
	f.saveCacheToFile()
}

func (f *FileCache) Display() {
	fmt.Println(f.cache)
}

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
