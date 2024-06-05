package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	bfs(root, &result, 0)
	reverse(&result)
	return result
}

func reverse(arr *[][]int) {
	i, j := 0, len(*arr)-1
	for i < j {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		i++
		j--
	}
}
