package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, word := range strs {
		charArray := []rune(word)
		sort.Slice(charArray, func(i, j int) bool {
			return charArray[i] < charArray[j]
		})
		sortedStr := string(charArray)
		if _, ok := strMap[sortedStr]; !ok {
			strMap[sortedStr] = make([]string, 0)
		}
		strMap[sortedStr] = append(strMap[sortedStr], word)
	}
	var result [][]string
	for _, val := range strMap {
		result = append(result, val)
	}
	return result
}
