package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	tests := []struct {
		in      []int
		workers int
		fn      func(int) int
		expect  []int
	}{{
		in:      []int{1, 2, 3, 4, 5},
		workers: 3,
		fn: func(n int) int {
			return n * 2
		},
		expect: []int{2, 4, 6, 8, 10},
	}, {
		in:      []int{1, 2, 3, 4, 5, 6},
		workers: 4,
		fn: func(n int) int {
			return n + 10
		},
		expect: []int{11, 12, 13, 14, 15, 16},
	}}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			if res := ConcurrentMap(tt.in, tt.workers, tt.fn); !slices.Equal(res, tt.expect) {
				t.Errorf("result: %v. expected: %v", res, tt.expect)
			}
		})
	}
}
