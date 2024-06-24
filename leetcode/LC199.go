package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		result = append(result, queue[n-1].Val)
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
	}
	return result
}
