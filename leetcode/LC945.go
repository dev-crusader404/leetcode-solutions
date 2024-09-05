package leetcode

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	var count int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			diff := nums[i-1] - nums[i] + 1
			nums[i] += diff
			count += diff
		}
	}
	return count
}
