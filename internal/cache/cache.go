package cache

import "sync"

type Cache struct {
	entries map[string]int
	mutex   *sync.RWMutex
}

func (c *Cache) Fetch(key string) int {
	c.mutex.RLock()
	value := c.entries[key]
	c.mutex.RUnlock()
	return value
}

func (c *Cache) Save(key string) {
	c.mutex.Lock()
	c.entries[key]++
	c.mutex.Unlock()
}

func New() *Cache {
	return &Cache{
		entries: make(map[string]int),
		mutex:   &sync.RWMutex{},
	}
}
