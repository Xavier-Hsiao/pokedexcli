package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte // raw data that we cache
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: map[string]cacheEntry{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.entries[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if entry, ok := cache.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

// Each time an interval passes, this function should remove any entries that are older than `interval`
// example: if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.
func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(interval)
	}
}

func (cache *Cache) reap(interval time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for key, entry := range cache.entries {
		if time.Since(entry.createdAt) > interval {
			delete(cache.entries, key)
		}
	}
}
