package leetcode

func longestContinuousSubstring(s string) int {
	count, maxCount := 1, 0
	if len(s) == 1 {
		return count
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i+1]-s[i] == 1 {
			count++
		} else {
			count = 1
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}
