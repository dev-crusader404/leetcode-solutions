package leetcode

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	min1, min2 := math.MaxInt, math.MaxInt
	if len(nums) < 3 {
		return false
	}

	for _, v := range nums {
		if v <= min1 {
			min1 = v
		} else if v <= min2 {
			min2 = v
		} else {
			return true
		}
	}
	return false
}

func RunLC334() {
	n := []int{2, 1, 5, 4, 3, 6}
	fmt.Println(increasingTriplet(n))
}
