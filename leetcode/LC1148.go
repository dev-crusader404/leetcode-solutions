package leetcode

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	var count int
	findGoodNode(root, &count, math.MinInt)
	return count
}

func findGoodNode(node *TreeNode, count *int, minNum int) {
	if node == nil {
		return
	}
	if node.Val >= minNum {
		*count++
		minNum = node.Val
	}
	findGoodNode(node.Left, count, minNum)
	findGoodNode(node.Right, count, minNum)
}

func RunLC1148() {
	x := []any{3, 1, 4, 3, nil, 1, 5}
	t := BuildTree(x)
	fmt.Println(goodNodes(t))
}
