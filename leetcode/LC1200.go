package leetcode

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minDiff := math.MaxInt
	var result [][]int
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < minDiff {
			minDiff = arr[i] - arr[i-1]
			result = result[:0]
			result = append(result, []int{arr[i-1], arr[i]})
		} else if arr[i]-arr[i-1] == minDiff {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}
