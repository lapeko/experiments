package main

import (
	"math/rand"
	"runtime"
	"testing"
)

func TestCalcSumConcurently(t *testing.T) {
	slice := builSlice(SLICE_SIZE, SLICE_ITEM_MAX_VALUE)

	if res, exp := calcSumConcurently(slice, int64(rand.Intn(runtime.NumCPU()*2))), sliceSum(slice); exp != res {
		t.Errorf("Expected: %d, Received: %d", exp, res)
	}
}

func BenchmarkSum(b *testing.B) {
	b.ReportAllocs()

	slice := builSlice(SLICE_SIZE, SLICE_ITEM_MAX_VALUE)

	for b.Loop() {
		sliceSum(slice)
	}
}

func BenchmarkCalcSumConcurently(b *testing.B) {
	b.ReportAllocs()

	slice := builSlice(SLICE_SIZE, SLICE_ITEM_MAX_VALUE)

	for b.Loop() {
		calcSumConcurently(slice, int64(rand.Intn(runtime.NumCPU()*2)))
	}
}
