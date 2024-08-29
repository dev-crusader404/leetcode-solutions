package leetcode

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	arr1 := make([]int, 26)
	arr2 := make([]int, 26)
	for _, v := range word1 {
		arr1[v-'a']++
	}
	for _, v := range word2 {
		arr2[v-'a']++
	}
	for i := 0; i < 26; i++ {
		if arr1[i] != 0 && arr2[i] == 0 || arr1[i] == 0 && arr2[i] != 0 {
			return false
		}
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
