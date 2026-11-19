package main

import (
	"runtime"
	"strings"
	"testing"
)

const (
	testLine          = "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
	internalLoopCount = 100
	metricName        = "total-alloc-MB"
)

func BenchmarkConcat(b *testing.B) {
	var begin, end runtime.MemStats

	runtime.ReadMemStats(&begin)
	b.ReportAllocs()

	for b.Loop() {
		res := ""
		for range internalLoopCount {
			res += testLine
		}
		_ = res
	}

	runtime.ReadMemStats(&end)
	mb := float64(end.TotalAlloc-begin.TotalAlloc) / 1024 / 1024
	b.ReportMetric(mb, metricName)
}

func BenchmarkStringBuilderGrow(b *testing.B) {
	var begin, end runtime.MemStats

	runtime.ReadMemStats(&begin)
	b.ReportAllocs()

	for b.Loop() {
		var res strings.Builder
		for range internalLoopCount {
			res.WriteString(testLine)
		}
	}

	runtime.ReadMemStats(&end)
	mb := float64(end.TotalAlloc-begin.TotalAlloc) / 1024 / 1024
	b.ReportMetric(mb, metricName)
}

func BenchmarkStringBuilder(b *testing.B) {
	var begin, end runtime.MemStats

	runtime.ReadMemStats(&begin)
	b.ReportAllocs()

	size := internalLoopCount * len(testLine) * b.N

	var res strings.Builder
	res.Grow(size)

	for b.Loop() {
		for range internalLoopCount {
			res.WriteString(testLine)
		}
	}

	runtime.ReadMemStats(&end)
	mb := float64(end.TotalAlloc-begin.TotalAlloc) / 1024 / 1024
	b.ReportMetric(mb, metricName)
}
