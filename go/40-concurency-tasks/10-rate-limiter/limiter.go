package main

import (
	"context"
	"sync"
	"time"
)

type rateLimiter struct {
	mu     sync.Mutex
	tokens int
}

func NewRateLimiter(size int) *rateLimiter {
	return &rateLimiter{
		tokens: size,
	}
}

func (r *rateLimiter) Wait(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.tokens == 0 {
		return overLimitError
	}

	r.tokens--
	timer := time.NewTimer(time.Second)

	go func() {
		defer func() {
			r.mu.Lock()
			defer r.mu.Unlock()
			timer.Stop()
			r.tokens++
		}()

		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			return
		}
	}()

	return nil
}
