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
	var wg sync.WaitGroup
	size := len(slice)
	blockSize := float64(size) / float64(workersCount)

	prevIdx := 0
	for i := range workersCount {
		idx := int(blockSize * float64(i+1))
		pIdx := prevIdx

		wg.Go(func() {
			sum.Add(sliceSum(slice[pIdx:idx]))
		})

		prevIdx = idx
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
