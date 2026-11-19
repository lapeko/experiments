package main

import (
	"fmt"
	"sync"
)

type MpmcQueue[T any] struct {
	ch chan T
}

func NewMpmcQueue[T any](size int) *MpmcQueue[T] {
	return &MpmcQueue[T]{
		ch: make(chan T, size),
	}
}

func (m *MpmcQueue[T]) Push(v T) {
	m.ch <- v
}

func (m *MpmcQueue[T]) Pull() (T, bool) {
	v, ok := <-m.ch
	return v, ok
}

func (m *MpmcQueue[T]) Close() {
	close(m.ch)
}

func main() {
	var produceWg sync.WaitGroup
	var receiveWg sync.WaitGroup

	queue := NewMpmcQueue[int](4)

	for i := range 2 {
		produceWg.Go(func() {
			for k := range 10 {
				queue.Push(i*10 + k)
			}
		})
	}

	go func() {
		produceWg.Wait()
		queue.Close()
	}()

	for range 3 {
		receiveWg.Go(func() {
			for {
				v, ok := queue.Pull()
				if !ok {
					return
				} else {
					fmt.Println(v)
				}
			}
		})
	}

	receiveWg.Wait()
}
