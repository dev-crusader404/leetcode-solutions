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
			rmLeft := s[:left] + s[left+1:]
			rmRight := s[:right] + s[right+1:]
			return checkPalindrome(rmLeft, 0, len(rmLeft)-1) || checkPalindrome(rmRight, 0, len(rmRight)-1)
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
