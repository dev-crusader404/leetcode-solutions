package leetcode

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		// right is sorted
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func FindMinInRotatedArray() {
	// 3,4,5,1,2
	// 4,5,6,1,2,3
	//4,5,1,2,3
	// 6,7,1,2,3,4,5
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findMin(arr))
}
