func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	var j int = 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i+1]
			j++
		}
	}
	return j
}