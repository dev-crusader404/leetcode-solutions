package leetcode

import "fmt"

func constrainedSubsetSum(nums []int, k int) int {
	res := nums[0]
	q := []int{}
	for i := range nums {
		if len(q) > 0 {
			nums[i] += q[0]
		}

		res = max(res, nums[i])
		for len(q) > 0 && q[len(q)-1] < nums[i] {
			q = q[:len(q)-1]
		}

		if nums[i] > 0 {
			q = append(q, nums[i])
		}

		if i >= k && len(q) > 0 && nums[i-k] == q[0] {
			q = q[1:]
		}
	}
	return res
}

func RunLC1425() {
	x := []int{-1, -2, -3}
	fmt.Println(constrainedSubsetSum(x, 2))
}
