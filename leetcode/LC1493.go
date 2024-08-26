package leetcode

func longestSubarray(nums []int) int {
	left, maxLen, zeroCount := 0, 0, 0

	for right, val := range nums {
		if val == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen - 1
}
