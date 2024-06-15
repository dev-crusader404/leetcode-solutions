package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	sumMap := make(map[int]int)
	sumMap[0] = 1
	return findPathSum(root, targetSum, 0, sumMap)
}

func findPathSum(node *TreeNode, target, currentSum int, m map[int]int) int {
	if node == nil {
		return 0
	}
	currentSum += node.Val
	var result int
	if v, ok := m[currentSum-target]; ok {
		result = v
	}
	m[currentSum]++
	result += findPathSum(node.Left, target, currentSum, m) + findPathSum(node.Right, target, currentSum, m)
	m[currentSum]--
	return result
}

func RunLC437() {
	a := []any{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1, 3, 5}
	t := BuildTree(a)
	fmt.Println(pathSum2(t, 8))
}
