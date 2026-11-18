package main

import (
	"fmt"
	"sync"
	"time"
)

type TTLCache interface {
	Set(key string, value any, ttl time.Duration)
	Get(key string) (any, bool)
	Delete(key string)
	Close()
}

type ttlItem struct {
	value any
	ttl   time.Time
}

type ttlCache struct {
	clean chan struct{}
	done  chan struct{}
	mu    sync.RWMutex
	m     map[string]ttlItem
	once  sync.Once
}

func NewTTLCache() TTLCache {
	c := &ttlCache{
		clean: make(chan struct{}),
		done:  make(chan struct{}),
		m:     make(map[string]ttlItem),
	}

	c.initCleaner()

	return c
}

func (c *ttlCache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = ttlItem{
		value: value,
		ttl:   time.Now().Add(ttl),
	}
}

func (c *ttlCache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.m[key]

	if !ok {
		return nil, false
	}

	if time.Now().After(v.ttl) {
		return nil, false
	}

	return v.value, true
}

func (c *ttlCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.m, key)
}

func (c *ttlCache) Close() {
	c.once.Do(func() {
		close(c.clean)
	})

	<-c.done
}

func (c *ttlCache) initCleaner() {
	t := time.NewTicker(time.Second)

	go func() {
		defer t.Stop()
		defer close(c.done)

		for {
			select {
			case <-c.clean:
				return
			case <-t.C:
				{
					func() {
						c.mu.Lock()
						defer c.mu.Unlock()

						now := time.Now()
						for k, v := range c.m {
							if now.After(v.ttl) {
								delete(c.m, k)
							}
						}
					}()
				}
			}
		}
	}()
}

func main() {
	c := NewTTLCache()
	defer func() {
		c.Close()
		fmt.Println("Closed")
	}()

	c.Set("a", 1, time.Second)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(c.Get("a"))
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(c.Get("a"))
}
