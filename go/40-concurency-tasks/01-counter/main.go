package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ROUTINES   = 100
	INCREMENTS = 1000
)

func runCounterAtomic() int {
	counter := atomic.Int64{}
	var wg sync.WaitGroup

	for range ROUTINES {
		wg.Go(func() {
			for range INCREMENTS {
				counter.Add(1)
			}
		})
	}

	wg.Wait()

	return int(counter.Load())
}

func runCounterMutex() int {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for range ROUTINES {
		wg.Go(func() {
			for range INCREMENTS {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		})
	}

	wg.Wait()

	return counter
}

func main() {
	start := time.Now()
	fmt.Println("Atomic counter: ", runCounterAtomic())
	atomicDuration := time.Since(start)
	fmt.Println("Duration", atomicDuration)

	start = time.Now()
	fmt.Println("Mutex counter: ", runCounterMutex())
	mutexDuration := time.Since(start)
	fmt.Println("Duration", mutexDuration)
	fmt.Printf("Atomic is %.2f times faster\n", float64(mutexDuration)/float64(atomicDuration))
}
