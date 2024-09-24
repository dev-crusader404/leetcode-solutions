package leetcode

func maximumCount(nums []int) int {
	var pos, neg int
	for i := range nums {
		if nums[i] > 0 {
			pos++
		} else if nums[i] < 0 {
			neg++
		}
	}
	return max(pos, neg)
}

func maximumCount2(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	var pos, neg int

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] < 0 {
			low = mid + 1
			neg = mid + 1
		} else if nums[mid] >= 0 {
			high = mid - 1
		}
	}
	low, high = 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] > 0 {
			high = mid - 1
			pos = n - mid
		} else if nums[mid] <= 0 {
			low = mid + 1
		}
	}
	return max(pos, neg)
}
