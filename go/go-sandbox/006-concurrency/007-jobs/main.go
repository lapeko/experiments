package main

import (
	"context"
	"fmt"
	"sync"
)

type jobData struct {
	i int
	v int
}

func Process(ctx context.Context, jobs []int, workers int) ([]int, error) {
	if workers < 1 {
		return jobs, fmt.Errorf("workers should be positive int")
	}

	var wg sync.WaitGroup
	in := make(chan jobData)
	out := make(chan jobData)

	go func() {
		defer close(in)
		for i, v := range jobs {
			select {
			case <-ctx.Done():
				return
			case in <- jobData{i: i, v: v}:
			}
		}
	}()

	for range workers {
		wg.Go(func() {
			for {
				select {
				case <-ctx.Done():
					return
				case j, ok := <-in:
					{
						if !ok {
							return
						} else {
							out <- jobData{i: j.i, v: j.v * 2}
						}
					}
				}
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	res := make([]int, len(jobs))
	for j := range out {
		select {
		case <-ctx.Done():
			return res, ctx.Err()
		default:
			res[j.i] = j.v
		}
	}
	return res, nil
}

func main() {
	jobs := []int{1, 2, 3, 4, 5}

	if result, err := Process(context.Background(), jobs, 3); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(result)
	}
}
