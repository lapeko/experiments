package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func randomWorker(ctx context.Context, out chan<- int, number int) {
	timer := time.NewTimer(time.Duration(rand.Int64N(1000)) * time.Millisecond)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		{
			fmt.Printf("Worker N: %d\texit\n", number)
			return
		}
	case <-timer.C:
		{
			n := rand.IntN(1000)
			fmt.Printf("Worker N: %d\tresult:%d\n", number, n)
			out <- n
		}

	}
}

func main() {
	var workerWg sync.WaitGroup
	var receiverWg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	out := make(chan int)

	go func() {
		time.Sleep(time.Duration(rand.Int64N(2000)) * time.Millisecond)
		cancel()
	}()

	for i := range 20 {
		workerWg.Go(func() {
			randomWorker(ctx, out, i+1)
		})
	}

	sum := 0
	receiverWg.Go(func() {
		for v := range out {
			sum += v
		}
	})

	workerWg.Wait()
	close(out)

	receiverWg.Wait()

	select {
	case <-ctx.Done():
		fmt.Println("Cancelled. Temp sum:", sum)
	default:
		fmt.Println(sum)
	}
}
