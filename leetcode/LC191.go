package leetcode

import "strconv"

func hammingWeight(n int) int {
	var c int
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			c++
		}
	}
	return c
}

func hammingWeight2(n int) int {
	var c int
	s := strconv.FormatInt(int64(n), 2)
	for _, v := range s {
		if v == '1' {
			c++
		}
	}
	return c
}
