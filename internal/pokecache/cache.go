package pokecache

import (
	// "fmt"
	"time"
	"sync"
)

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	Entries map[string]CacheEntry
	Mutex sync.Mutex
	Duration time.Duration
}

func NewCache(duration time.Duration) (cache *Cache) {
	cache = &Cache{
		Duration: duration,
		Entries: make(map[string]CacheEntry),
		Mutex: sync.Mutex{},
		}
	go cache.reapLoop()
	return
}

func (c *Cache) Add(key string, val []byte) {
	// fmt.Printf("  [[Adding cache entry for: %s]]\n", key)
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Entries[key] = CacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) (val []byte, exists bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// Had to declare func level [entry] to avoid masking
	var entry CacheEntry

	if entry, exists = c.Entries[key]; exists { 
		// fmt.Printf("  [[Retrieved entry for: %s]]\n", key)
		val = entry.val 
	} else {
		// fmt.Printf("  [[No entry found for: %s]]\n", key)
	}
	return
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Duration)
	defer ticker.Stop()
	for range ticker.C { c.reap() }
}

func (c *Cache) reap() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	for key, entry := range c.Entries {
		if time.Since(entry.createdAt) > c.Duration { 
			// fmt.Printf("\n  [[Reaping: %s]]\n", key)
			delete(c.Entries, key) 
		}
	}
}
