package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	sCopy := s[:3]
	s = append(s, 1, 2, 3)
	s = append(s, 4)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("sCopy: %v\n", sCopy)
}
