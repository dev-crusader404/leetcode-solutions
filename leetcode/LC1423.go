package leetcode

import (
	"fmt"
	"math"
)

// K-window solution
func maxScore2(cardPoints []int, k int) int {
	var totalSum int

	for i := 0; i < k; i++ {
		totalSum += cardPoints[i]
	}
	maxPoints := totalSum
	left, right := k-1, len(cardPoints)-1

	for i := 0; i < k; i++ {
		totalSum = totalSum - cardPoints[left] + cardPoints[right]
		maxPoints = max(maxPoints, totalSum)
		left--
		right--
	}
	return maxPoints
}

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
