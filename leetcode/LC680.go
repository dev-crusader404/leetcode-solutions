package leetcode

func validPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	left, right := 0, len(s)-1

	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return checkPalindrome(s, left+1, right) || checkPalindrome(s, left, right-1)
		}
	}
	return true
}

func checkPalindrome(s string, L, R int) bool {
	for L < R {
		if s[L] != s[R] {
			return false
		}
		L++
		R--
	}
	return true
}
