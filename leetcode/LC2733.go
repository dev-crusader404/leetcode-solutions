package leetcode

func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	mxm := max(nums[0], nums[1])
	mim := min(nums[0], nums[1])
	n := nums[2]

	if mim < n && n < mxm {
		return n
	} else if mim > n {
		return mim
	}
	return mxm
}
