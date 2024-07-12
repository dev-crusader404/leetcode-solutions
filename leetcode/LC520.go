package leetcode

import (
	"strings"
)

func detectCapitalUse(word string) bool {
	caps := strings.ToUpper(word)
	lower := strings.ToLower(word)
	FirstCap := strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	return caps == word || lower == word || FirstCap == word
}
