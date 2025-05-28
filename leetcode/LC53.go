package leetcode

import "fmt"

func maxSubArray(nums []int) int {
	maxTillHere, maxSoFar := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxTillHere = max(maxTillHere+nums[i], nums[i])
		maxSoFar = max(maxSoFar, maxTillHere)
	}
	return maxSoFar
}

func RunLC53() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}
