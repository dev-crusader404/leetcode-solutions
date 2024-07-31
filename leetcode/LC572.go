package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(node *TreeNode, sub *TreeNode) bool {
	if node == nil || sub == nil {
		return node == sub
	}

	if node.Val != sub.Val {
		return false
	}
	return isSameTree(node.Left, sub.Left) && isSameTree(node.Right, sub.Right)
}
