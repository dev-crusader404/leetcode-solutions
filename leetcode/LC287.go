package leetcode

//Solution with O(N) space and O(N) time complexity
func findDuplicate(nums []int) int {
	idx := 0
	for idx < len(nums) {
		if nums[idx] != nums[nums[idx]-1] {
			nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
		} else {
			idx++
		}
	}

	for i, v := range nums {
		if i+1 != v {
			return v
		}
	}
	return nums[0]
}

func RunLC287() {
	n := []int{3, 3, 3, 3, 3}
	println(findDuplicate(n))
}
