package main

import (
	"sync"
)

type cache struct {
	m    map[string]any
	rwmu sync.RWMutex
}

func New() *cache {
	return &cache{m: make(map[string]any)}
}

func (c *cache) Set(key string, value any) {
	c.rwmu.Lock()
	defer c.rwmu.Unlock()

	c.m[key] = value
}

func (c *cache) Get(key string) (any, bool) {
	c.rwmu.RLock()
	defer c.rwmu.RUnlock()

	v, ok := c.m[key]
	return v, ok
}

func (c *cache) Delete(key string) {
	c.rwmu.Lock()
	defer c.rwmu.Unlock()

	delete(c.m, key)
}
