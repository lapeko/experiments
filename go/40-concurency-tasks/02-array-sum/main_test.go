package main

import (
	"math/rand"
	"runtime"
	"testing"
)

var (
	testSlice = builSlice(SLICE_SIZE, SLICE_ITEM_MAX_VALUE)
	workers   = int64(rand.Intn(runtime.NumCPU() * 2))
)

func TestCalcSumConcurently(t *testing.T) {
	slice := builSlice(SLICE_SIZE, SLICE_ITEM_MAX_VALUE)

	if res, exp := calcSumConcurently(slice, workers), sliceSum(slice); exp != res {
		t.Errorf("Expected: %d, Received: %d", exp, res)
	}
}

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		sliceSum(testSlice)
	}
}

func BenchmarkCalcSumConcurently(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		calcSumConcurently(testSlice, workers)
	}
}
