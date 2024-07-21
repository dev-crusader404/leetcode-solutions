package leetcode

import "math"

func minimumDeletions(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	maxm, minm := math.MinInt, math.MaxInt
	var minIdx, maxIdx int
	for i, v := range nums {
		if v > maxm {
			maxm = v
			maxIdx = i
		}

		if v < minm {
			minm = v
			minIdx = i
		}
	}
	fromLeft := max(minIdx, maxIdx) + 1
	fromRight := len(nums) - min(minIdx, maxIdx)
	var removal int
	if minIdx > maxIdx {
		removal = (maxIdx + 1) + len(nums) - minIdx
	} else {
		removal = (minIdx + 1) + len(nums) - maxIdx
	}
	return min(fromLeft, min(fromRight, removal))
}
