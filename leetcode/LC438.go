package leetcode

import "fmt"

/*

438. Find All Anagrams in a String
Solved
Medium
Topics
Companies
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".


Constraints:

1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.

*/

func findAnagrams(s string, p string) []int {

	if len(s) < len(p) {
		return []int{}
	}

	var result []int
	cMap := make(map[rune]int)

	for _, v := range p {
		cMap[v]++
	}
	var left int
	missingCount := len(p)

	for right, val := range s {
		if right-left+1 > len(p) {
			if v, ok := cMap[rune(s[left])]; ok {
				if v >= 0 {
					missingCount++
				}
				cMap[rune(s[left])]++
			}
			left++
		}

		if v, ok := cMap[val]; ok {
			if v > 0 {
				missingCount--
			}
			cMap[val]--
		}

		if right-left+1 == len(p) {
			if missingCount == 0 {
				result = append(result, left)
			}
		}
	}
	return result
}

func RunFindAnagram() {
	s := "abcbaabca"
	p := "aabc"

	r := findAnagrams(s, p)
	fmt.Println(r)
}
