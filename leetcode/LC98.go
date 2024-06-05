package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validateBST(root, math.MinInt, math.MaxInt)
}

func validateBST(node *TreeNode, minimum, maximum int) bool {
	if node == nil {
		return true
	}

	if minimum < node.Val && node.Val < maximum {
		return validateBST(node.Left, minimum, node.Val) && validateBST(node.Right, node.Val, maximum)
	} else {
		return false
	}
}
