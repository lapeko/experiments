package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	WORKERS_NUM  = 5
	IN_SIZE      = 20
	MAX_RAND_INT = 1000
)

func multiplyByTenWorker(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 10
	}
}

func runNWorkers(n int, in <-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	if n < 1 {
		n = 1
	}

	for range n {
		wg.Go(func() {
			multiplyByTenWorker(in, out)
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := make(chan int)

	out := runNWorkers(WORKERS_NUM, in)

	go func() {
		defer close(in)
		for range IN_SIZE {
			in <- rand.Intn(MAX_RAND_INT)
		}
	}()

	for v := range out {
		fmt.Println(v)
	}
}
