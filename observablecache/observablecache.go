package observablecache

type LocalCache struct {
	store map[string]string
}

func (c *LocalCache) Get(key string) string {
	return c.store[key]
}

func (c *LocalCache) Set(key string, value string) {
	c.store[key] = value
}

func New() LocalCache {
	return LocalCache{
		store: make(map[string]string),
	}
}
