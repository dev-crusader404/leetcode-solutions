package leetcode

import "fmt"

func mostCompetitive(nums []int, k int) []int {
	stk := []int{}
	for i := 0; i < len(nums); i++ {
		for len(stk) > 0 && stk[len(stk)-1] > nums[i] && (len(nums)-i > k-len(stk)) {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 || len(stk) < k {
			stk = append(stk, nums[i])
		}
	}
	return stk
}

func RunLC1673() {
	a := []int{6, 8, 6, 2, 5, 6, 4, 2, 5, 7, 1, 5, 8, 7}
	fmt.Println(mostCompetitive(a, 4))
}
