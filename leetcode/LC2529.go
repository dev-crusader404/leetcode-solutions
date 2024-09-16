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
