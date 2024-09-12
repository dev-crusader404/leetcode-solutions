package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	currLvlSum := root.Val

	for len(queue) > 0 {
		n := len(queue)
		nextLvlSum := 0
		for i := 0; i < n; i++ {
			node := queue[i]
			node.Val = currLvlSum - node.Val
			if node.Left != nil {
				nextLvlSum += node.Left.Val
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				nextLvlSum += node.Right.Val
				queue = append(queue, node.Right)
			}

			if node.Left != nil && node.Right != nil {
				sum := node.Left.Val + node.Right.Val
				node.Left.Val = sum
				node.Right.Val = sum
			}
		}
		queue = queue[n:]
		currLvlSum = nextLvlSum
	}
	return root
}
