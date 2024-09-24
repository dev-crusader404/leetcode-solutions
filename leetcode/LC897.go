package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	return increasingTree(root, nil)
}

func increasingTree(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}
	res := increasingTree(root.Left, root)
	root.Left = nil
	root.Right = increasingTree(root.Right, tail)
	return res
}
