package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newCacheEntery := cacheEntry{}
	newCacheEntery.Val = value
	newCacheEntery.createdAt = time.Now()

	c.Entries[key] = newCacheEntery
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cacheEntry, ok := c.Entries[key]
	if !ok {
		return []byte{}, false
	}

	return cacheEntry.Val, true
}

func (c *Cache) reapLoop(lifespan time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	currentTime := time.Now()
	for key, entry := range c.Entries {
		timeDiff := currentTime.Sub(entry.createdAt)
		if timeDiff >= lifespan {
			delete(c.Entries, key)
		}
	}

}

func NewCache(reapInterval time.Duration) Cache {
	cache := Cache{
		Entries: map[string]cacheEntry{},
		mu:      &sync.RWMutex{},
	}

	ticker := time.NewTicker(reapInterval)

	go func() {
		for range ticker.C {
			cache.reapLoop(reapInterval)
		}
	}()

	return cache
}
