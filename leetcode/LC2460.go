package leetcode

func applyOperations(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}

		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	return nums
}
