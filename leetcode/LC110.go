package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return findBalance(root) != -1
}

func findBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := findBalance(node.Left)
	right := findBalance(node.Right)
	if left == -1 || right == -1 {
		return -1
	}
	if abs2(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func abs2(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
