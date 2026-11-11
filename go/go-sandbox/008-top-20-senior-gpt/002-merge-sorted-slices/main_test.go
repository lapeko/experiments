package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestMergeSortedSlices(t *testing.T) {
	tests := []struct {
		s1  []int
		s2  []int
		res []int
	}{
		{
			s1:  []int{1, 6, 8},
			s2:  []int{2, 7},
			res: []int{1, 2, 6, 7, 8},
		},
		{
			s1:  []int{},
			s2:  []int{},
			res: []int{},
		},
		{
			s1:  []int{1, 1},
			s2:  []int{1},
			res: []int{1, 1, 1},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			if res := mergeSortedSlices(tt.s1, tt.s2); !slices.Equal(res, tt.res) {
				t.Errorf("received %v when %v expected", res, tt.res)
			}
		})
	}
}
