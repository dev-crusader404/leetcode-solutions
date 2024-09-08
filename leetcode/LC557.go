package leetcode

import "strings"

func reverseWords557(s string) string {
	words := strings.Fields(s)
	for i, val := range words {
		str := reverseString2(val)
		words[i] = str
	}
	return strings.Join(words, " ")
}

func reverseString2(s string) string {
	word := []byte(s)
	l, r := 0, len(s)-1
	for l < r {
		word[l], word[r] = word[r], word[l]
		l++
		r--
	}
	return string(word)
}
