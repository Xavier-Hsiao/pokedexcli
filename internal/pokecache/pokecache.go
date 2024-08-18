package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time // To purge to cache
	val       []byte
}

func CreateNewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	// Should do the reapLoop in a saparate goroutine
	// otherwise `c` will never get returned
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEntry, ok := c.cache[key]
	return cacheEntry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	// Returns a new Ticker containing a channel
	// that will send the current time on the channel after each tick
	ticker := time.NewTicker(interval)
	for range ticker.C {
		now := time.Now()
		for key, entry := range c.cache {
			if now.Sub(entry.createdAt) > interval {
				delete(c.cache, key)
			}
		}
	}
}
