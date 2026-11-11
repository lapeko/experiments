package main

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		expected []int
	}{
		{
			name:     "Case 1",
			in:       []int{1, 2, 4, 6, 2, 1, 3},
			expected: []int{1, 2, 4, 6, 3},
		},
		{
			name:     "Case 2 [empty]",
			in:       []int{},
			expected: []int{},
		},
		{
			name:     "Case 3",
			in:       []int{1, 1},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := RemoveDuplicates(test.in); !slices.Equal(res, test.expected) {
				t.Errorf("Received %v when %v expected", res, test.expected)
			}
		})
	}
}
