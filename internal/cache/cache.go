package cache

import "sync"

type cache struct {
	entries map[string]int
	mutex   *sync.RWMutex
}

func (c *cache) Fetch(key string) int {
	c.mutex.RLock()
	value := c.entries[key]
	c.mutex.RUnlock()
	return value
}

func (c *cache) Save(key string) {
	c.mutex.Lock()
	c.entries[key]++
	c.mutex.Unlock()
}

func New() *cache {
	return &cache{
		entries: make(map[string]int),
		mutex:   &sync.RWMutex{},
	}
}
