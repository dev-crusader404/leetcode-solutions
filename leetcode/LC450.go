package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	root = bstDelete(root, key)
	return root
}

func bstDelete(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return node
	}

	if key < node.Val {
		node.Left = bstDelete(node.Left, key)
	} else if key > node.Val {
		node.Right = bstDelete(node.Right, key)
	} else if key == node.Val {
		if node.Left == nil {
			node = node.Right
		} else if node.Right == nil {
			node = node.Left
		} else {
			node.Val = getMin(node.Right)
			node.Right = bstDelete(node.Right, node.Val)
		}
	}
	return node
}

func getMin(node *TreeNode) int {
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}
