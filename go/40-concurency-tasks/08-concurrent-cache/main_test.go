package main

import (
	"math/rand"
	"sync"
	"testing"
)

const (
	CONCURENCY_NUM = 1_000
	KEY            = "a"
)

func TestCache(t *testing.T) {
	var wg sync.WaitGroup
	c := New()

	wg.Go(func() {
		for range CONCURENCY_NUM {
			c.Set(KEY, rand.Int())
		}
	})

	for range 10 {
		wg.Go(func() {
			for range CONCURENCY_NUM {
				c.Get(KEY)
			}
		})
	}

	wg.Go(func() {
		for range CONCURENCY_NUM {
			c.Delete(KEY)
		}
	})

	wg.Wait()
	c.Set(KEY, CONCURENCY_NUM)
	if v, ok := c.Get(KEY); v != CONCURENCY_NUM || !ok {
		t.Errorf("result: %v. expected: %d", v, CONCURENCY_NUM)
	}
}
