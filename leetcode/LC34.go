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

func searchRange2(nums []int, target int) []int {
	defValue := []int{-1, -1}

	if len(nums) == 0 {
		return defValue
	}

	low, high := 0, len(nums)-1
	index := binarySearch(low, high, target, nums)
	if index == -1 {
		return defValue
	}
	rightIndex, leftIndex := index, index
	var temp int
	for leftIndex != -1 {
		temp = leftIndex
		leftIndex = binarySearch(low, leftIndex-1, target, nums)
	}
	leftIndex = temp

	for rightIndex != -1 {
		temp = rightIndex
		rightIndex = binarySearch(rightIndex+1, high, target, nums)
	}
	rightIndex = temp

	return []int{leftIndex, rightIndex}
}

func binarySearch(low, high, target int, nums []int) int {
	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func RunLC34() {
	arr := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange2(arr, 5))
}
