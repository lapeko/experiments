package main

import "fmt"

func groupAnagrams(lines []string) [][]string {
	m := make(map[[26]int][]string)
	res := [][]string{}

	for _, line := range lines {
		key := getKey(line)
		m[key] = append(m[key], line)
	}

	for _, lines := range m {
		res = append(res, lines)
	}

	return res
}

var (
	aKey = 'a'
	zKey = 'z'
)

func getKey(line string) [26]int {
	var res [26]int

	for _, key := range line {
		if key < aKey || key > zKey {
			continue
		}
		res[key-aKey]++
	}

	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"abba", "baab", "baba", "aab", "baa"}))
}
