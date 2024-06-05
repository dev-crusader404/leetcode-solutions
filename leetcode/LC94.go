package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return []int{}
	}
	inorder(root, &res)
	return res
}

func inorder(node *TreeNode, result *[]int) {
	if node.Left != nil {
		inorder(node.Left, result)
	}
	*result = append((*result), node.Val)
	if node.Right != nil {
		inorder(node.Right, result)
	}
}
