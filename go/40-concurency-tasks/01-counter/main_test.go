package main

import "testing"

func BenchmarkRunCounterAtomic(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		runCounterAtomic()
	}
}

func BenchmarkRunCounterMutex(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		runCounterMutex()
	}
}
