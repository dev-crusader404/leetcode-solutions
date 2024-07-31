package leetcode

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	var i, j int
	for i < len(word1) && j < len(word2) {
		sb.WriteByte(word1[i])
		sb.WriteByte(word2[j])
		i++
		j++
	}
	if i < len(word1) {
		sb.WriteString(word1[i:])
	}
	if j < len(word2) {
		sb.WriteString(word2[j:])
	}
	return sb.String()
}
