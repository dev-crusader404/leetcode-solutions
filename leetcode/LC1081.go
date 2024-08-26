package leetcode

import "sort"

func heightChecker(heights []int) int {
	dup := make([]int, len(heights))
	copy(dup, heights)
	sort.Ints(heights)
	wrIdx := 0
	for i, v := range heights {
		if v != dup[i] {
			wrIdx++
		}
	}
	return wrIdx
}
