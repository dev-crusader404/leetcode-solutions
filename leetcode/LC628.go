package leetcode

import "sort"

func maximumProduct2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	m1 := nums[0] * nums[1] * nums[n-1]
	m2 := nums[n-1] * nums[n-2] * nums[n-3]
	return max(m1, m2)
}
