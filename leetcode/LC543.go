package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	len := 0
	findMaxDiameter(root, &len)
	return len
}

func findMaxDiameter(node *TreeNode, l *int) int {
	if node == nil {
		return 0
	}
	left := findMaxDiameter(node.Left, l)
	right := findMaxDiameter(node.Right, l)
	*l = max(*l, left+right)
	return max(left, right) + 1
}
