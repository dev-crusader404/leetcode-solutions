package leetcode

import "strings"

func findOcurrences(text string, first string, second string) []string {
	splitWord := strings.Split(text, " ")
	result := []string{}
	if len(splitWord) < 3 {
		return result
	}

	for i, w := range splitWord {
		if i > 0 && i < len(splitWord)-1 && splitWord[i-1] == first && w == second {
			result = append(result, splitWord[i+1])
		}
	}
	return result
}
