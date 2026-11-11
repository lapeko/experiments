// CODE EXAMPLE VALID FOR COMPILING
package main

import (
	"fmt"
	"sync"
)

func multiplyByIndex(nums []int) []int {
	res := make([]int, len(nums))
	var wg sync.WaitGroup

	for i, v := range nums {
		wg.Go(func() {
			res[i] = i * v
		})
	}

	wg.Wait()

	return res
}

func main() {
	fmt.Println(multiplyByIndex([]int{2, 3, 5, 7})) // [0 3 10 21]
}
