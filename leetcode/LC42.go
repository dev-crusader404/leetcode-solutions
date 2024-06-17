package leetcode

import "fmt"

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	var total int
	for i := 0; i < len(height); i++ {
		maxL, maxR := height[i], height[i]
		left, right := i, i
		for left >= 0 {
			maxL = max(maxL, height[left])
			left--
		}
		for right < len(height) {
			maxR = max(maxR, height[right])
			right++
		}
		total += min(maxL, maxR) - height[i]
	}
	return total
}

func RunLC42() {
	h := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(h))
}
