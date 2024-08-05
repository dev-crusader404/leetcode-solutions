package leetcode

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	prefix[0], suffix[len(nums)-1] = 1, 1
	for i := range nums {
		if i > 0 {
			prefix[i] = prefix[i-1] * nums[i-1]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-1 {
			suffix[i] = suffix[i+1] * nums[i+1]
		}
	}
	for i := range nums {
		nums[i] = prefix[i] * suffix[i]
	}
	return nums
}
