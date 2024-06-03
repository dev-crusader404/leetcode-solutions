package leetcode

import (
	"fmt"
	"sort"
)

func merge56(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	idx := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[idx][1] < intervals[i][0] {
			idx++
			intervals[idx] = intervals[i]
		} else {
			intervals[idx][1] = max(intervals[idx][1], intervals[i][1])
		}
	}
	return intervals[:idx+1]
}

func RunLC56() {
	arr := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge56(arr))
}
