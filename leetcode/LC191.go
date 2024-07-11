package leetcode

import "strconv"

func hammingWeight(n int) int {
	var c int
	s := strconv.FormatInt(int64(n), 2)
	for _, v := range s {
		if v == '1' {
			c++
		}
	}
	return c
}
