package leetcode

import (
	"fmt"
	"reflect"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1, s2 := []int{}, []int{}
	dfs(root1, &s1)
	dfs(root2, &s2)
	return reflect.DeepEqual(s1, s2)
}

func dfs(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*arr = append(*arr, node.Val)
	}
	dfs(node.Left, arr)
	dfs(node.Right, arr)
}

func RunLC872() {
	t1 := []any{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}
	t2 := []any{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}
	tree1 := BuildTree(t1)
	tree2 := BuildTree(t2)
	fmt.Println(leafSimilar(tree1, tree2))
}
