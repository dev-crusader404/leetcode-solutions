package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return findSmallString(root, "")
}

func findSmallString(node *TreeNode, result string) string {
	result = string(node.Val+97) + result
	if node.Left == nil && node.Right == nil {
		return result
	}
	var left, right string
	if node.Right != nil {
		right = findSmallString(node.Right, result)
	}
	if node.Left != nil {
		left = findSmallString(node.Left, result)
	}

	if node.Left == nil {
		return right
	} else if node.Right == nil {
		return left
	} else {
		return min(left, right)
	}
}
