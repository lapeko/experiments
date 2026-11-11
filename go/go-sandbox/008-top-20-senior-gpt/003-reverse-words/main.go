package main

import "fmt"

func reverseWords(text string) string {
	var words [][]rune
	word := []rune{}
	spaceRune := rune(' ')

	for _, ch := range text {
		if ch == spaceRune {
			if len(word) > 0 {
				words = append(words, word)
			}
			word = []rune{}
		} else {
			word = append(word, ch)
		}
	}
	if len(word) > 0 {
		words = append(words, word)
	}

	var res []rune
	for i := len(words) - 1; i >= 0; i-- {
		res = append(res, words[i]...)
		if i != 0 {
			res = append(res, rune(' '))
		}
	}

	return string(res)
}

func main() {
	fmt.Println(reverseWords("Раз два"))
}
