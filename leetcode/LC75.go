package leetcode

import "fmt"

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	i, low, high := 0, 0, len(nums)-1

	for i <= high {
		if nums[i] == 0 {
			nums[low], nums[i] = nums[i], nums[low]
			low++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[high], nums[i] = nums[i], nums[high]
			high--
		}
	}
	fmt.Println(nums)
}
