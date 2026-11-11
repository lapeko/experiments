package main

func RemoveDuplicates(slice []int) []int {
	set, res := make(map[int]struct{}), []int{}

	for _, v := range slice {
		if _, ok := set[v]; !ok {
			res = append(res, v)
			set[v] = struct{}{}
		}
	}

	return res
}

func main() {

}
