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
	id1, isRight := crossHalf(minIdx, len(nums))
	id2, isRight2 := crossHalf(maxIdx, len(nums))

	if isRight == isRight2 {
		return max(id1, id2)
	}
	return id1 + id2
}

func crossHalf(i, n int) (int, bool) {
	if i >= n/2 {
		return (n - i), true
	}
	return i + 1, false
}
