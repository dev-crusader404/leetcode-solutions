package leetcode

import (
	"strings"
)

func detectCapitalUse(word string) bool {
	var uCount, lCount int
	for _, v := range word {
		if v >= 'A' && v <= 'Z' {
			uCount++
		} else if v >= 'a' && v <= 'z' {
			lCount++
		} else if uCount > 1 && lCount > 1 {
			return false
		}
	}

	if len(word) == uCount || len(word) == lCount || (uCount == 1 && word[0] >= 'A' && word[0] <= 'Z') {
		return true
	}
	return false
}

func detectCapitalUse2(word string) bool {
	caps := strings.ToUpper(word)
	lower := strings.ToLower(word)
	FirstCap := strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	return caps == word || lower == word || FirstCap == word
}
