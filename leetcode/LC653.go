package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	return findElement(root, k, m)
}

func findElement(node *TreeNode, k int, m map[int]struct{}) bool {
	if node == nil {
		return false
	}
	if _, ok := m[k-node.Val]; ok {
		return true
	}
	m[node.Val] = struct{}{}
	return findElement(node.Left, k, m) || findElement(node.Right, k, m)
}
