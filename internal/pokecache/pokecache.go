package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data  map[string]cacheEntry
	mutex sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]cacheEntry),
	}
	if ttl > 0 {
		go cache.reapLoop(ttl)
	}
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, exists := c.data[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(ttl time.Duration) {
	ticker := time.NewTicker(ttl)
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > ttl {
				delete(c.data, key)
			}
		}
		c.mutex.Unlock()
	}
	ticker.Stop()
}
