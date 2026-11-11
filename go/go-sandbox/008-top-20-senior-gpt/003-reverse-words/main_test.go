package main

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		in     string
		expect string
	}{{
		in:     "Раз два",
		expect: "два Раз",
	}, {
		in:     "One two three",
		expect: "three two One",
	}}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			if res := reverseWords(tt.in); res != tt.expect {
				t.Errorf("Received %s when %s expected", res, tt.expect)
			}
		})
	}
}
