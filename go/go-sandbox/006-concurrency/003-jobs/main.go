package main

import (
	"fmt"
	"sync"
)

type worker struct {
	jobs    <-chan int
	results chan<- int
}

func (w *worker) run() {
	for n := range w.jobs {
		w.results <- n * 2
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)
	var workerGroup sync.WaitGroup
	var consumerGroup sync.WaitGroup

	for range 3 {
		w := worker{
			jobs:    jobs,
			results: results,
		}
		workerGroup.Go(func() {
			w.run()
		})
	}

	go func() {
		for i := range 10 {
			jobs <- i + 1
		}
		close(jobs)
	}()

	sum := 0
	consumerGroup.Go(func() {
		for r := range results {
			sum += r
		}
	})

	workerGroup.Wait()
	close(results)
	consumerGroup.Wait()
	fmt.Println(sum)
}
