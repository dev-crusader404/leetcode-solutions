package leetcode

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) < 2 {
		return intervals
	}

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := make([][]int, 0)
	index := 0
	for index < len(intervals) && intervals[index][1] < newInterval[0] {
		result = append(result, intervals[index])
		index++
	}

	for index < len(intervals) && intervals[index][0] <= newInterval[1] {
		newInterval[0] = min(intervals[index][0], newInterval[0])
		newInterval[1] = max(intervals[index][1], newInterval[1])
		index++
	}
	result = append(result, newInterval)

	for index < len(intervals) && intervals[index][0] > newInterval[1] {
		result = append(result, intervals[index])
		index++
	}
	return result
}
