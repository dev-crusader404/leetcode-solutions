package leetcode

func minimumAverageDifference(nums []int) int {
	var totalSum, prefixSum int
	for _, v := range nums {
		totalSum += v
	}

	minDiff, index := totalSum/(len(nums)), len(nums)-1

	for i := 0; i < len(nums)-1; i++ {
		prefixSum += nums[i]
		avgDiff := absolute((prefixSum)/(i+1) - (totalSum-prefixSum)/(len(nums)-i-1))
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
