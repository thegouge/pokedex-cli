package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cacheEntry, ok := c.entries[key]

	return cacheEntry.val, ok
}

func (c *Cache) reapLoop(lifespan time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	currentTime := time.Now()
	ticker := time.NewTicker(lifespan)

	for range ticker.C {
		for key, entry := range c.entries {
			timeDiff := currentTime.Sub(entry.createdAt)
			if timeDiff >= lifespan {
				delete(c.entries, key)
			}
		}
	}

}

func NewCache(reapInterval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.RWMutex{},
	}

	go cache.reapLoop(reapInterval)

	return cache
}
