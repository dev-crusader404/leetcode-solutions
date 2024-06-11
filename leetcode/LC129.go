package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return calculateSum(root, 0)
}

func calculateSum(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	sum = sum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return sum
	}
	return calculateSum(node.Left, sum) + calculateSum(node.Right, sum)
}
