package leetcode

func minOperations1827(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var oprn int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			diff := nums[i-1] - nums[i] + 1
			oprn += diff
			nums[i] += diff
		}
	}
	return oprn
}
