package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	q := []*TreeNode{root}
	res := make([]float64, 0)
	for len(q) > 0 {
		n := len(q)
		var sum float64
		for i := 0; i < n; i++ {
			node := q[i]
			sum += float64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, sum/float64(n))
		q = q[n:]
	}
	return res
}
