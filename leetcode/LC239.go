package leetcode

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	stk := []int{}
	res := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		for len(stk) > 0 && i-stk[0]+1 > k {
			stk = stk[1:]
		}
		for len(stk) > 0 && nums[stk[len(stk)-1]] < v {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)

		if i+1 >= k {
			res = append(res, nums[stk[0]])
		}
	}
	return res
}

func RunLC239() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(a, 1))
}
