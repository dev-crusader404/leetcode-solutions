package leetcode

func sortedSquares(nums []int) []int {
	if len(nums) < 1 {
		return []int{}
	}
	var result []int = make([]int, len(nums))
	left, right, index := 0, len(nums)-1, len(nums)-1

	for left <= right {
		if (nums[left] * nums[left]) < (nums[right] * nums[right]) {
			result[index] = nums[right] * nums[right]
			right--
		} else {
			result[index] = nums[left] * nums[left]
			left++
		}
		index--
	}
	return result
}
