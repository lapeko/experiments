package main

import (
	"reflect"
	"strconv"
	"testing"
)

func AssertFunction[T1, T2 any](t *testing.T, name string, given []T1, expected []T2, fn func(v T1) T2) {
	t.Run(name, func(t *testing.T) {
		res := Map(given, fn)
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("Return %v when %v expected", res, expected)
		}
	})
}

func TestMap(t *testing.T) {
	AssertFunction(t, "Test 1",
		[]int{1, 2, 3},
		[]int{2, 4, 6},
		func(v int) int {
			return v * 2
		},
	)

	AssertFunction(t, "Test 2",
		[]int{1, 2, 3},
		[]string{"1", "2", "3"},
		func(v int) string {
			return strconv.Itoa(v)
		},
	)
}
