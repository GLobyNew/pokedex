package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	store map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	newC := &Cache{
		store: make(map[string]cacheEntry),
		mu:    sync.Mutex{},
	}
	go newC.reapLoop(interval)
	return newC
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cEntry, exist := c.store[key]; exist {
		return cEntry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.store {
			if time.Since(entry.createdAt) > interval {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}
}
