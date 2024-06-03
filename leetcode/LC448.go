package leetcode

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 1 {
		return []int{}
	}
	idx := 0
	result := make([]int, 0)

	for idx < len(nums) {
		if nums[idx] != nums[nums[idx]-1] {
			nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
		} else {
			idx++
		}
	}

	for i, v := range nums {
		if i+1 != v {
			result = append(result, i+1)
		}
	}
	return result
}
