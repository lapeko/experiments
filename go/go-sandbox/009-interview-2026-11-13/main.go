package main

import (
	"fmt"
	"math"
)

func findSimpeNumbersSum(n int) []int {
	res := []int{}

	if n < 2 {
		return res
	}

	res = append(res, 2)

	for v := range n - 2 {
		v = v + 2
		met := false
		for _, vNatural := range res {
			r := float64(v) / float64(vNatural)
			if r == math.Floor(r) {
				met = true
				break
			}
		}
		if !met {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	fmt.Println(findSimpeNumbersSum(100))
}
