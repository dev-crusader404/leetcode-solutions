package leetcode

import (
	"fmt"
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	l := []int{}
	for i := left; i <= right; i++ {
		if selfDivide(i) {
			l = append(l, i)
		}
	}
	return l
}

func selfDivide(x int) bool {
	for i := x; i > 0; i /= 10 {
		num := i % 10
		if i == 0 || x%num != 0 {
			return false
		}
	}
	return true
}

func selfDivide2(x int) bool {
	s := strconv.Itoa(x)
	for _, v := range s {
		num := int(v - '0')
		if num == 0 || x%num != 0 {
			return false
		}
	}
	return true
}

func RunLC728() {
	fmt.Println(selfDividingNumbers(1, 22))
}
