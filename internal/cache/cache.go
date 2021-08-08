package cache

type Cache struct {
	Entries map[string]int
}

func (c *Cache) Fetch(key string) int {
	value := c.Entries[key]
	return value
}

func (c *Cache) Save(key string) {
	c.Entries[key]++
}

func New() *Cache {
	return &Cache{
		Entries: make(map[string]int),
	}
}
