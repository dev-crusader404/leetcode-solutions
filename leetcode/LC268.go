package leetcode

func missingNumber(nums []int) int {
	idx := 0
	for idx < len(nums) {
		if nums[idx] < len(nums) && nums[idx] != nums[nums[idx]] {
			nums[idx], nums[nums[idx]] = nums[nums[idx]], nums[idx]
		} else {
			idx++
		}
	}

	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}
