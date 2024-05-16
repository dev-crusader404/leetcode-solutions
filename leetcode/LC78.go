package leetcode

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	c := []int{}
	backtrack(&result, c, nums, 0)
	return result
}

func backtrack(result *[][]int, current, nums []int, start int) {
	copiedArray := make([]int, len(current))
	copy(copiedArray, current)
	*result = append(*result, copiedArray)

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		backtrack(result, current, nums, i+1)
		current = current[:len(current)-1]
	}
}

func Subset78() {
	arr := []int{1, 2, 3}
	fmt.Println(subsets(arr))
}
