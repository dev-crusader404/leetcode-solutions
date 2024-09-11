package leetcode

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	lvl, maxSum, mxLvl := 0, math.MinInt, 0
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		lvl++
		for i := 0; i < n; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if sum > maxSum {
			maxSum = sum
			mxLvl = lvl
		}
		queue = queue[n:]
	}
	return mxLvl
}

func RunLC1161() {
	x := []any{1, 7, 0, 7, -8, nil, nil}
	t := BuildTree(x)
	fmt.Println(maxLevelSum(t))
}
