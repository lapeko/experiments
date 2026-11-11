package main

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		name     string
		payload  []int
		expected []int
	}{
		{
			name:     "Shold pass",
			payload:  []int{1, 2, 3, 1, 2},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Shold fail",
			payload:  []int{1, 2, 3, 1, 2},
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			res := removeDuplicates(test.payload)
			if !slices.Equal(res, test.expected) {
				t.Errorf("Expected: %v, received: %v", test.expected, res)
			}
		})
	}
}

func TestFindSkipped(t *testing.T) {
	tests := []struct {
		name     string
		payload  []int
		expected int
	}{
		{
			name:     "Shoudl pass",
			payload:  []int{1, 2, 4, 5},
			expected: 3,
		},
		{
			name:     "Shoudl pass",
			payload:  []int{1, 2, 3, 4, 5, 7},
			expected: 6,
		},
		{
			name:     "Shoudl pass",
			payload:  []int{1, 2, 3, 4, 5, 6, 8, 9},
			expected: 7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if res := findSkipped(test.payload); res != test.expected {
				t.Errorf("Result: %d when %d expected", res, test.expected)
			}
		})
	}
}
