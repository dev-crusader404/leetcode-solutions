package leetcode

/*

Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.



Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false


Constraints:

1 <= s1.length, s2.length <= 104
s1 and s2 consist of lowercase English letters.

*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	cmap1 := make(map[rune]int)
	for _, v := range s1 {
		cmap1[v]++
	}
	var left int
	cmap2 := make(map[rune]int)
	for right, val := range s2 {
		cmap2[val]++

		if (right - left + 1) > len(s1) {
			cmap2[rune(s2[left])]--
			if cmap2[rune(s2[left])] == 0 {
				delete(cmap2, rune(s2[left]))
			}
			left++
		}
		if checkPermutation(cmap1, cmap2) {
			return true
		}
	}
	return false
}

func checkPermutation(c1, c2 map[rune]int) bool {
	if len(c1) > len(c2) {
		return false
	}
	for k, v := range c1 {
		var v2 int
		var ok bool
		if v2, ok = c2[k]; !ok {
			return false
		}
		if v2 != v {
			return false
		}
	}
	return true
}
