package leetcode

import (
	"math"
)

func thirdMax(nums []int) int {
	m1, m2, m3 := math.MinInt, math.MinInt, math.MinInt
	for _, val := range nums {
		if val >= m1 {
			if val != m1 {
				m3 = m2
				m2 = m1
				m1 = val
			}
		} else if val >= m2 {
			if val != m2 {
				m3 = m2
				m2 = val
			}
		} else if val >= m3 {
			m3 = val
		}
	}
	if m3 == math.MinInt {
		return m1
	}
	return m3
}
