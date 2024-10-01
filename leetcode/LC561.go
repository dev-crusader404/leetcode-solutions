package leetcode

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var sum int
	for i := 0; i < len(nums); i = i + 2 {
		sum += nums[i]
	}
	return sum
}
