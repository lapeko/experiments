package main

func mergeSortedSlices(s1, s2 []int) []int {
	ptr1, ptr2 := 0, 0
	len1, len2 := len(s1), len(s2)
	res := make([]int, len1+len2)

	for ptr1 < len1 || ptr2 < len2 {
		ptr := ptr1 + ptr2
		moveLeft := true

		switch {
		case ptr1 < len1 && ptr2 < len2:
			if s1[ptr1] > s2[ptr2] {
				moveLeft = false
			}
		case ptr1 >= len1:
			moveLeft = false
		}

		if moveLeft {
			res[ptr] = s1[ptr1]
			ptr1++
		} else {
			res[ptr] = s2[ptr2]
			ptr2++
		}
	}

	return res
}

func main() {

}
