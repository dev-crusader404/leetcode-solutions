package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findDepth(root, 0)
}

func findDepth(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}
	depth++
	return max(findDepth(node.Left, depth), findDepth(node.Right, depth))
}
