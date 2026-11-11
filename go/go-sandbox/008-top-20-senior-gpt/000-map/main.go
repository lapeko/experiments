package main

import (
	"fmt"
	"strconv"
)

func Map[T1 any, T2 any](slice []T1, fn func(v T1) T2) []T2 {
	res := make([]T2, len(slice))

	for i, v := range slice {
		res[i] = fn(v)
	}

	return res
}

func main() {
	// ints := []int{1, 2, 3, 4, 5}
	var ints []int

	fmt.Printf("%v\n", Map(ints, func(d int) string {
		return strconv.Itoa(d)
	}))
}
