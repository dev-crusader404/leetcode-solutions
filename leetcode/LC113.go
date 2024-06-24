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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	findSum(root, targetSum, []int{}, &result)
	return result
}

func findSum(node *TreeNode, target int, arr []int, result *[][]int) {
	if node == nil {
		return
	}

	arr = append(arr, node.Val)
	if node.Val == target && node.Left == nil && node.Right == nil {
		dup := make([]int, len(arr))
		copy(dup, arr)
		(*result) = append((*result), dup)
		return
	}

	if node.Left != nil {
		findSum(node.Left, target-node.Val, arr, result)
	}

	if node.Right != nil {
		findSum(node.Right, target-node.Val, arr, result)
	}
}

func RunLC113() {
	arr := []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}
	tree := BuildTree(arr)
	result := pathSum(tree, 22)
	fmt.Println(result)
}
