package leetcode

import "fmt"

func minDeletion(nums []int) int {
	var numDeletion int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] && (i-numDeletion)%2 == 0 {
			numDeletion++
		}
	}
	return numDeletion + (len(nums)-numDeletion)%2
}

func RunLC2216() {
	// a := []int{1, 1, 2, 3, 5}
	a := []int{1, 1, 1, 1, 2, 2, 3, 3}

	fmt.Println(minDeletion(a))
}
