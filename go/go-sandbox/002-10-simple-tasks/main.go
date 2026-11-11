package main

func removeDuplicates(slice []int) []int {
	m := make(map[int]bool)
	res := []int{}
	for _, val := range slice {
		if ok := m[val]; !ok {
			res = append(res, val)
			m[val] = true
		}
	}
	return res
}

func findSkipped(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	expectedSum := (1 + (len(slice) + 1)) * (len(slice) + 1) / 2
	return expectedSum - sum
}
