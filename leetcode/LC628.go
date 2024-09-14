package leetcode

import (
	"math"
	"sort"
)

func maximumProduct2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	m1 := nums[0] * nums[1] * nums[n-1]
	m2 := nums[n-1] * nums[n-2] * nums[n-3]
	return max(m1, m2)
}

func maximumProduct3(nums []int) int {
	m1, m2, m3 := math.MinInt, math.MinInt, math.MinInt
	n1, n2 := math.MaxInt, math.MaxInt

	for i := range nums {
		if nums[i] > m1 {
			m3 = m2
			m2 = m1
			m1 = nums[i]
		} else if nums[i] > m2 {
			m3 = m2
			m2 = nums[i]
		} else if nums[i] > m3 {
			m3 = nums[i]
		}

		if nums[i] < n1 {
			n2 = n1
			n1 = nums[i]
		} else if nums[i] < n2 {
			n2 = nums[i]
		}
	}
	return max(n1*n2*m1, m1*m2*m3)
}
