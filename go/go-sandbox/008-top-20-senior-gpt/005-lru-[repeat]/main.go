package main

import "container/list"

type LRU[T any] interface {
	Get(string) (T, bool)
	Put(string, T)
}

type lru[T any] struct {
	size int
	m    map[string]*list.Element
	dll  *list.List
}

type listValue[T any] struct {
	key string
	val T
}

func NewLRU[T any](size int) LRU[T] {
	return &lru[T]{
		size: size,
		m:    make(map[string]*list.Element),
		dll:  list.New(),
	}
}

func (l *lru[T]) Get(key string) (T, bool) {
	el, ok := l.m[key]
	var zero T
	if !ok {
		return zero, false
	}
	l.dll.MoveToFront(el)
	v, _ := el.Value.(listValue[T])
	return v.val, true
}

func (l *lru[T]) Put(key string, val T) {
	newListVal := listValue[T]{
		key: key,
		val: val,
	}
	el, ok := l.m[key]

	if !ok {
		if l.dll.Len() >= l.size {
			last := l.dll.Remove(l.dll.Back())
			lastListVal, _ := last.(listValue[T])
			delete(l.m, lastListVal.key)
		}
		l.m[key] = l.dll.PushFront(newListVal)
		return
	}

	el.Value = newListVal
	l.dll.MoveToFront(el)
}
