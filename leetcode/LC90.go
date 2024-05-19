package leetcode

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	backtrackDup(&result, nums, []int{}, 0)
	return result
}

func backtrackDup(result *[][]int, arr, current []int, start int) {
	dupArr := make([]int, len(current))
	copy(dupArr, current)
	*result = append(*result, dupArr)
	for i := start; i < len(arr); i++ {
		if i > start && arr[i] == arr[i-1] {
			continue
		}
		current = append(current, arr[i])
		backtrackDup(result, arr, current, i+1)
		current = current[:len(current)-1]
	}
}

func Subset90() {
	// [[],[1],[1,2],[1,2,2],[2],[2,2]]
	arr := []int{3, 1, 2, 2, 2}
	fmt.Println(subsetsWithDup(arr))
}
