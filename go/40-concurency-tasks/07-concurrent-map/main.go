package main

import "sync"

type chanPayload struct {
	index int
	value int
}

func ConcurrentMap(
	input []int,
	workers int,
	fn func(int) int,
) []int {
	ch := make(chan chanPayload)
	res := make([]int, len(input))
	var wg sync.WaitGroup

	for range workers {
		wg.Go(func() {
			for pl := range ch {
				res[pl.index] = fn(pl.value)
			}
		})
	}

	for i, v := range input {
		ch <- chanPayload{i, v}
	}
	close(ch)

	wg.Wait()

	return res
}
