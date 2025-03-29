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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest2(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var min int
	return findKSmallest(root, &k, min)
}

func findKSmallest(node *TreeNode, k *int, min int) int {
	if node.Left != nil {
		min = findKSmallest(node.Left, k, min)
	}
	(*k)--
	if (*k) == 0 {
		min = node.Val
		return min
	}
	if node.Right != nil {
		min = findKSmallest(node.Right, k, min)
	}
	return min
}
