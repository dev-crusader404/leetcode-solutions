package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return sum
			}

			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
