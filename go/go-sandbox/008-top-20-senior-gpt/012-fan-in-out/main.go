package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fanIn[T any](ctx context.Context, ins ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	for _, in := range ins {
		wg.Go(func() {
			for v := range in {
				select {
				case <-ctx.Done():
					return
				case out <- v:
				}
			}
		})
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func sleepRandom(ctx context.Context, maxMs int) {
	d := time.Duration(rand.Intn(maxMs)) * time.Microsecond
	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case <-ctx.Done():
		return
	case <-t.C:
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	in1 := make(chan int)
	in2 := make(chan int)

	out := fanIn(ctx, in1, in2)

	wg.Go(func() {
		defer close(in1)
		for i := range 10 {
			select {
			case <-ctx.Done():
				return
			case in1 <- i + 1:
				sleepRandom(ctx, 500)
			}
		}
	})

	wg.Go(func() {
		defer close(in2)
		for i := range 10 {
			select {
			case <-ctx.Done():
				return
			case in2 <- (i + 1) * 1000:
				sleepRandom(ctx, 500)
			}
		}
	})

	wg.Go(func() {
		for v := range out {
			fmt.Println(v)
		}
	})

	wg.Wait()
}
