package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	stk := []*TreeNode{}

	for _, v := range nums {
		node := &TreeNode{Val: v}
		for len(stk) > 0 && stk[len(stk)-1].Val < v {
			n := len(stk) - 1
			node.Left = stk[n]
			stk = stk[:n]
		}
		if len(stk) > 0 {
			stk[len(stk)-1].Right = node
		}
		stk = append(stk, node)
	}

	if len(stk) == 0 {
		return nil
	}
	return stk[len(stk)-1]
}

func RunLC654() {
	arr := []int{3, 2, 1, 6, 0, 5}
	fmt.Println(constructMaximumBinaryTree(arr))
}
