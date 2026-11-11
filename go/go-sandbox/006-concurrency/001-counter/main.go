package main

import (
	"fmt"
	"sync"
)

// func concurrentCounter() int {
// 	conter := atomic.Int64{}
// 	var wg sync.WaitGroup
//
// 	for range 100 {
// 		wg.Go(func() {
// 			conter.Add(1000)
// 		})
// 	}
//
// 	wg.Wait()
// 	return int(conter.Load())
// }

func concurrentCounter() int {
	counter := 0
	var wg sync.WaitGroup
	var mux sync.Mutex

	for range 100 {
		wg.Go(func() {
			mux.Lock()
			defer mux.Unlock()

			counter += 1000
		})
	}

	wg.Wait()
	return counter
}

func main() {
	fmt.Println(concurrentCounter())
}
