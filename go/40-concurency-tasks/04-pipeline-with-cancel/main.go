package main

import "fmt"

func genInts(count int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range count {
			out <- i + 1
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func print(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	print(square(genInts(100)))
}
