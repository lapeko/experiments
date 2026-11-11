package main

import (
	"fmt"
	"sync"
)

func runner(n int, ch chan<- (int)) {
	begin, end := n*10, n*10+10
	for i := begin; i < end; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan (int))
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Go(func() {
			runner(i, ch)
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
