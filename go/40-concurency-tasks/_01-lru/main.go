package main

import (
	"container/list"
	"sync"
)

type LRU[T any] interface {
	Get(string) (T, bool)
	Set(string, T)
}

type elementValue[T any] struct {
	k string
	v T
}

type lru[T any] struct {
	size int
	mu   sync.Mutex
	m    map[string]*list.Element
	l    *list.List
}

func NewLRU[T any](size int) LRU[T] {
	if size < 1 {
		panic("size should be positive int")
	}
	return &lru[T]{
		size: size,
		m:    make(map[string]*list.Element),
		l:    list.New(),
	}
}

func (l *lru[T]) Get(k string) (T, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	el, ok := l.m[k]
	if !ok {
		var zero T
		return zero, false
	}

	l.l.MoveToFront(el)

	return el.Value.(elementValue[T]).v, true
}

func (l *lru[T]) Set(k string, v T) {
	l.mu.Lock()
	defer l.mu.Unlock()

	el, ok := l.m[k]
	if ok {
		newVal := el.Value.(elementValue[T])
		newVal.v = v
		el.Value = newVal
		l.l.MoveToFront(el)
		return
	}

	e := l.l.PushFront(elementValue[T]{k: k, v: v})
	if len(l.m) == l.size {
		ev := l.l.Remove(l.l.Back())
		k := ev.(elementValue[T]).k
		delete(l.m, k)
	}
	l.m[k] = e
}
