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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	height := findHeight(root)
	if height == 0 {
		return 1
	}

	upperTreeCount := int(math.Pow(2, float64(height))) - 1
	low, high := 0, upperTreeCount

	for low < high {
		idxToFind := int(math.Ceil(float64(low+high) / 2))
		if nodeExist(root, idxToFind, height) {
			low = idxToFind
		} else {
			high = idxToFind - 1
		}
	}
	leafCount := low + 1
	return leafCount + upperTreeCount
}

func nodeExist(node *TreeNode, idx, height int) bool {
	low, high, count := 0, int(math.Pow(2, float64(height)))-1, 0

	for count < height {
		mid := int(math.Ceil(float64(low+high) / 2))
		if idx >= mid {
			node = node.Right
			low = mid
		} else {
			node = node.Left
			high = mid - 1
		}
		count++
	}
	return node != nil
}

func findHeight(node *TreeNode) int {
	var h int
	for node.Left != nil {
		h++
		node = node.Left
	}
	return h
}

func RunLC222() {
	arr := []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	t := BuildTree(arr)
	fmt.Println(countNodes(t))
}
