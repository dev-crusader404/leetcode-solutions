package leetcode

func pivotIndex(nums []int) int {
	var leftSum, rightSum int

	for _, n := range nums {
		leftSum += n
	}

	for i := 0; i < len(nums); i++ {
		leftSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		rightSum += nums[i]
	}
	return -1
}
