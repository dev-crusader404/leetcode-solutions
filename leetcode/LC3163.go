package leetcode

import (
	"strconv"
	"strings"
)

func compressedString(word string) string {
	var sb strings.Builder
	right := 0
	for right < len(word) {
		left := right
		for right < len(word) && word[left] == word[right] && (right-left) < 9 {
			right++
		}
		sb.WriteString(strconv.Itoa(right - left))
		sb.WriteByte(word[left])
	}
	return sb.String()
}
