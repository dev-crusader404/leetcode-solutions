package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight, rightHeight := minDepth(root.Left), minDepth(root.Right)

	if leftHeight == 0 || rightHeight == 0 {
		return leftHeight + rightHeight + 1
	}

	return 1 + min(leftHeight, rightHeight)
}
