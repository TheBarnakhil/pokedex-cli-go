package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		interval: interval,
		entry:    map[string]cacheEntry{},
	}

	go cache.ReapLoopCache()

	return cache
}
