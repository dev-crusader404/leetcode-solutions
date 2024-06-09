package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	preorder(root, &result)
	return result
}

func preorder(node *TreeNode, res *[]int) {
	*res = append((*res), node.Val)
	if node.Left != nil {
		preorder(node.Left, res)
	}

	if node.Right != nil {
		preorder(node.Right, res)
	}
}
