package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flattenII(root *TreeNode) {
	if root == nil {
		return
	}
	node := root
	for node != nil {
		if node.Left != nil {
			maxChild := node.Left
			for maxChild.Right != nil {
				maxChild = maxChild.Right
			}
			maxChild.Right = node.Right
			node.Right = node.Left
			node.Left = nil
		}
		node = node.Right
	}
}
