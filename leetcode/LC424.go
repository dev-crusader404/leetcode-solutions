package leetcode

/*
Given a string s and an integer k, return the length of the longest substring of s such that the frequency of each character in this substring is greater than or equal to k.
if no such substring exists, return 0.

Example 1:

Input: s = "aaabb", k = 3
Output: 3
Explanation: The longest substring is "aaa", as 'a' is repeated 3 times.
Example 2:

Input: s = "ababbc", k = 2
Output: 5
Explanation: The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.


Constraints:

1 <= s.length <= 104
s consists of only lowercase English letters.
1 <= k <= 105
*/

import (
	"math"
)

func CharacterReplacement(s string, k int) int {
	if k > len(s) {
		return len(s)
	}
	cMap := make(map[rune]int)

	var left, maxCount int
	maxLength := math.MinInt32

	for right, val := range s {
		cMap[val]++
		currentLength := right - left + 1
		maxCount = int(math.Max(float64(maxCount), float64(cMap[val])))

		for currentLength-maxCount > k {
			cMap[rune(s[left])]--
			left++
			currentLength = right - left + 1
		}

		if maxLength < currentLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
