package leetcode

func isSubsequence(s string, t string) bool {
	var left, right int
	if len(s) > len(t) {
		return false
	}
	for left < len(s) && right < len(t) {
		if s[left] == t[right] {
			left++
		}
		right++
	}
	return left == len(s)
}
