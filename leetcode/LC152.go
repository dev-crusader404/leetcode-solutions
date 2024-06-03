package leetcode

import "math"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums) - 1
	left, right, maxVal := 0, 0, math.MinInt

	for i := 0; i <= n; i++ {
		left = nonZeroVal(left) * nums[i]
		right = nonZeroVal(right) * nums[n-i]
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

func nonZeroVal(x int) int {
	if x == 0 {
		return 1
	}
	return x
}
