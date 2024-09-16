package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var foundX, foundY bool
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Val == x {
				foundX = true
			}
			if node.Val == y {
				foundY = true
			}
			if node.Left != nil && node.Right != nil {
				if node.Left.Val == x && node.Right.Val == y {
					return false
				}
				if node.Left.Val == y && node.Right.Val == x {
					return false
				}
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		if foundX && foundY {
			return true
		}
	}
	return false
}
