package leetcode

import "math"

func getMinDistance(nums []int, target int, start int) int {
	var minIndex int = math.MaxInt
	for i, v := range nums {
		if v == target {
			minIndex = min(minIndex, pos(i-start))
		}
	}
	return minIndex
}

func pos(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
