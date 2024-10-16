package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	sum := 0
	getTilt(root, &sum)
	return sum
}

func getTilt(node *TreeNode, sum *int) int {
	var left, right int
	if node == nil {
		return 0
	}
	if node.Left != nil {
		left = getTilt(node.Left, sum)
	}
	if node.Right != nil {
		right = getTilt(node.Right, sum)
	}
	diff := left - right
	if diff < 0 {
		diff *= -1
	}
	*sum += diff
	return left + right + node.Val
}
