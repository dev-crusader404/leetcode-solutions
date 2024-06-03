package leetcode

func findDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	idx := 0
	for idx < len(nums) {
		if nums[idx] != nums[nums[idx]-1] {
			nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
		} else {
			idx++
		}
	}
	result := make([]int, 0)
	for i, v := range nums {
		if i+1 != v {
			result = append(result, v)
		}
	}
	return result
}
