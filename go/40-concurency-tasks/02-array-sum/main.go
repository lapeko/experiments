package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

const (
	SLICE_SIZE           = 1_000_000
	WORKERS_SIZE         = 14
	SLICE_ITEM_MAX_VALUE = 1_000_000
)

func builSlice(size, maxItemValue int64) []int64 {
	slice := make([]int64, size)
	for i := range slice {
		slice[i] = rand.Int63n(maxItemValue)
	}
	return slice
}

func calcSumConcurently(slice []int64, workersCount int64) int64 {
	if workersCount < 1 {
		return sliceSum(slice)
	}
	sum := atomic.Int64{}
	size := int64(len(slice))
	var wg sync.WaitGroup

	for i := range workersCount {
		start := i * size / workersCount
		end := (i + 1) * size / workersCount

		wg.Go(func() {
			sum.Add(sliceSum(slice[start:end]))
		})
	}

	wg.Wait()

	return sum.Load()
}

func sliceSum(slice []int64) int64 {
	var sum int64
	for _, v := range slice {
		sum += v
	}
	return sum
}

func main() {
	slice := builSlice(SLICE_SIZE, SLICE_ITEM_MAX_VALUE)

	concurentSum := calcSumConcurently(slice, WORKERS_SIZE)
	notConcurentSum := sliceSum(slice)

	fmt.Printf("Received: %d, Expected: %d. Correct: %t\n", concurentSum, notConcurentSum, notConcurentSum == concurentSum)
}
