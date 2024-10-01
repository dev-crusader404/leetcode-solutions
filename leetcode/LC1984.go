package leetcode

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	var left int
	minDiff := math.MaxInt
	for right := 0; right < len(nums); right++ {
		if right-left+1 == k {
			minDiff = min(minDiff, nums[right]-nums[left])
			left++
		}
	}
	return minDiff
}
