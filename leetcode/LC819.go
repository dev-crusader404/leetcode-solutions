package leetcode

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	m := make(map[string]int)
	ban := make(map[string]struct{})
	for _, b := range banned {
		ban[b] = struct{}{}
	}
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(paragraph, f)
	for _, w := range words {
		if _, exist := ban[w]; !exist {
			m[w]++
		}
	}
	maxim, result := 0, ""
	for k, v := range m {
		if v > maxim {
			maxim = v
			result = k
		}
	}
	return result
}
