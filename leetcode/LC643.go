package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var maxAvg float64 = -math.MaxFloat64
	var left, sum int
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		if right-left == k-1 {
			maxAvg = max(maxAvg, float64(sum)/float64(k))
			sum -= nums[left]
			left++
		}
	}
	return maxAvg
}
