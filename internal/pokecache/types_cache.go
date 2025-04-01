package pokecache

import (
	"sync"
	"time"
)

type (
	cacheEntry struct {
		createdAt time.Time
		val       []byte
	}

	Cache struct {
		entry    map[string]cacheEntry
		interval time.Duration
		mu       sync.RWMutex
	}
)
