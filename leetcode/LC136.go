package leetcode

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var n int
	for _, v := range nums {
		n = n ^ v
	}
	return n
}
