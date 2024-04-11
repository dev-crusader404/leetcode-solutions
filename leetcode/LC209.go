package leetcode

import "math"

func MinSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minLength := math.MaxInt32
	leftIndex := 0
	sum := 0
	for i, v := range nums {
		sum += v
		for sum >= target {
			if minLength > (i - leftIndex + 1) {
				minLength = i - leftIndex + 1
			}
			sum -= nums[leftIndex]
			leftIndex++
		}
	}

	if minLength == math.MaxInt32 {
		minLength = 0
	}
	return minLength
}
