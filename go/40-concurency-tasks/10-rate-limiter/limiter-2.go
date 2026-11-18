package main

import (
	"container/list"
	"context"
	"sync"
	"time"
)

type rateLimiter2 struct {
	mu          sync.Mutex
	limit       int
	timesWindow list.List
}

func NewRateLimiter2(size int) *rateLimiter2 {
	return &rateLimiter2{
		limit: size,
	}
}

func (r *rateLimiter2) Wait(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handleWindow()

	if r.timesWindow.Len() >= r.limit {
		return overLimitError
	}

	r.timesWindow.PushBack(time.Now())
	return nil
}

func (r *rateLimiter2) handleWindow() {
	for r.timesWindow.Len() > 0 && time.Since(r.timesWindow.Front().Value.(time.Time)) >= time.Second {
		r.timesWindow.Remove(r.timesWindow.Front())
	}
}
