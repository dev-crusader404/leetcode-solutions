package leetcode

import "fmt"

func nextGreaterElements(nums []int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		index := (i + 1) % len(nums)
		for nums[i] >= nums[index] && i != index {
			index++
			if index == len(nums) {
				index %= len(nums)
			}
		}
		if index == i {
			result = append(result, -1)
		} else {
			result = append(result, nums[index])
		}
	}
	return result
}

func RunLC503() {
	a := []int{1, 5, 4, 3, 2, 1}
	fmt.Println(nextGreaterElements(a))
}
