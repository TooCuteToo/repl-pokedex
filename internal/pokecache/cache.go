package pokecache

import (
	"fmt"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		interval: interval,
		data:     make(map[string]cacheEntry),
	}

	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()
		for t := range ticker.C {
			fmt.Println("reapLoop at: ", t)
			c.mu.Lock()
			c.reapLoop()
			c.mu.Unlock()
		}
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return ce.val, true
}

func (c *Cache) reapLoop() {
	for k, v := range c.data {
		if time.Now().Sub(v.createdAt) >= c.interval {
			delete(c.data, k)
		}
	}
}
