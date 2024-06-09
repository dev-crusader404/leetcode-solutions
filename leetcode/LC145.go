package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	postorder(root, &result)
	return result
}

func postorder(node *TreeNode, res *[]int) {
	if node.Left != nil {
		postorder(node.Left, res)
	}

	if node.Right != nil {
		postorder(node.Right, res)
	}

	*res = append((*res), node.Val)
}
