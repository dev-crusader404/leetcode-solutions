package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	even := true
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		list := []int{}
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if even {
				list = append(list, node.Val)
			} else {
				list = append([]int{node.Val}, list...)
			}
		}
		queue = queue[l:]
		even = !even
		result = append(result, list)
	}
	return result
}
