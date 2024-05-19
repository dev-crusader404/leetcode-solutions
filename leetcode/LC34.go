package leetcode

import "fmt"

func searchRange(nums []int, target int) []int {
	defValue := []int{-1, -1}

	if len(nums) == 0 {
		return defValue
	}

	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return findboundary(mid, nums)
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return defValue
}

func findboundary(mid int, nums []int) []int {
	target := nums[mid]
	i, j := mid, mid
	for i > 0 && nums[i-1] == target {
		i--
	}

	for j < len(nums)-1 && nums[j+1] == target {
		j++
	}
	return []int{i, j}
}

func RunLC34() {
	arr := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(arr, 5))
}
