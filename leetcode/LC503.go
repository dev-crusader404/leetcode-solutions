package leetcode

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	for i, _ := range result {
		result[i] = -1
	}
	stack := make([]int, 0)
	for i := 0; i < 2*len(nums); i++ {
		for len(stack) > 0 && nums[i%n] > nums[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return result
}

func nextGreaterElements2(nums []int) []int {
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
	fmt.Println(nextGreaterElements2(a))
}
