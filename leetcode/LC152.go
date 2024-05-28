package leetcode

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums) - 1
	left, right, maxVal := 1, 1, math.MinInt

	for i := 0; i <= n; i++ {
		left *= nums[i]
		right *= nums[n-i]
		maxVal = getMax(maxVal, getMax(left, right))
	}
	return maxVal
}

func getMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
