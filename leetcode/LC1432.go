package leetcode

import (
	"fmt"
	"math"
)

func maxScore(cardPoints []int, k int) int {
	var totalSum, prefixSum, left int
	maxPoints := math.MinInt
	for _, v := range cardPoints {
		totalSum += v
	}
	if len(cardPoints) == k {
		return totalSum
	}

	for right, val := range cardPoints {
		prefixSum += val
		if (right - left + 1) == len(cardPoints)-k {
			maxPoints = max(maxPoints, totalSum-prefixSum)
			prefixSum -= cardPoints[left]
			left++
		}
	}
	return maxPoints
}

func RunLC1432() {
	p := []int{9, 7, 7, 9, 7, 7, 9}
	fmt.Println(maxScore(p, 7))
}
