package leetcode

import "math"

func reverses(x int) int {
	var result int
	for x != 0 {
		result = result*10 + (x % 10)
		x /= 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
