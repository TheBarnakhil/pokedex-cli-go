package pokecache

import (
	"time"
)

func (c *Cache) AddToCache(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) GetFromCache(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.entry[key]
	return entry.val, exists
}

func (c *Cache) deleteFromCache(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entry, key)
}

func (c *Cache) ReapLoopCache() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for t := range ticker.C {
		// ticker.Reset(0 * time.Second)
		for k, v := range c.entry {
			if t.Sub(v.createdAt) >= c.interval {
				c.deleteFromCache(k)
			}
		}
	}
}
