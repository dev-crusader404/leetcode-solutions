package leetcode

import "math"

func minimumAverageDifference(nums []int) int {
	var suffixSum, prefixSum int
	for _, v := range nums {
		suffixSum += v
	}

	minDiff, index, endAvg, n := math.MaxInt, 0, 0, len(nums)

	for i := 0; i < n; i++ {
		prefixSum += nums[i]
		suffixSum -= nums[i]
		if i == len(nums)-1 {
			endAvg = 0
		} else {
			endAvg = suffixSum / (n - i - 1)
		}
		avgDiff := absolute((prefixSum)/(i+1) - endAvg)
		if minDiff > avgDiff {
			minDiff = avgDiff
			index = i
		}
	}
	return index
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
