package leetcode

import "fmt"

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	left, right, mxArea := 0, len(height)-1, 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		mxArea = max(mxArea, area)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return mxArea
}

func RunLC11() {
	a := []int{1, 2, 4, 3}
	fmt.Println(maxArea(a))
}
