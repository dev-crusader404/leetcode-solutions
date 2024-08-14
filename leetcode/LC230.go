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

func kthSmallest(root *TreeNode, k int) int {
	q := []*TreeNode{root}

	for len(q) > 0 {
		for root != nil {
			q = append(q, root)
			root = root.Left
		}
		root = q[len(q)-1]
		q = q[:len(q)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return 0
}

func RunLC230() {
	arr := []any{5, 3, 6, 2, 4, nil, nil, 1}
	tree := BuildTree(arr)
	fmt.Println(kthSmallest(tree, 3))
}
