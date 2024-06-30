package leetcode

func isArraySpecial(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums)-1; i++ {
		sum = nums[i] + nums[i+1]
		if sum%2 == 0 {
			return false
		}
	}
	return true
}
