package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == root.Right {
		return -1
	}

	left := root.Left.Val
	right := root.Right.Val

	if left == root.Val {
		left = findSecondMinimumValue(root.Left)
	}
	if right == root.Val {
		right = findSecondMinimumValue(root.Right)
	}

	if left != -1 && right != -1 {
		return min(left, right)
	} else if left == -1 {
		return right
	} else {
		return left
	}
}
