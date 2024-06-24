package leetcode

import "fmt"

func trap2(height []int) int {
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

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	var maxL, maxR, total int
	left, right := 0, len(height)-1

	for left < right {
		if height[left] >= height[right] {
			if maxR < height[right] {
				maxR = height[right]
			} else {
				total += maxR - height[right]
			}
			right--
		} else {
			if maxL < height[left] {
				maxL = height[left]
			} else {
				total += maxL - height[left]
			}
			left++
		}
	}
	return total
}

func RunLC42() {
	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(h))
}
