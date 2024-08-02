package leetcode

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	var maxLen int = math.MinInt
	start := 0
	for i := range s {
		findPalindrome(s, i, i, &maxLen, &start)
		findPalindrome(s, i, i+1, &maxLen, &start)
	}
	return s[start : start+maxLen]
}

func findPalindrome(s string, l, r int, m, start *int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	if (*m) < (r - l - 1) {
		*m = r - l - 1
		*start = l + 1
	}
}

func RunLC5() {
	fmt.Println(longestPalindrome("babac"))
}
