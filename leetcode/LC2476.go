package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodes(root *TreeNode, queries []int) [][]int {
	sortedArr := []int{}
	dfsTransverse(root, &sortedArr)
	var result [][]int
	for _, n := range queries {
		res := binarySearchNode(sortedArr, n)
		result = append(result, res)
	}
	return result
}

func binarySearchNode(arr []int, n int) []int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == n {
			return []int{arr[mid], arr[mid]}
		} else if arr[mid] < n {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if start == len(arr) {
		return []int{arr[end], -1}
	} else if end == -1 {
		return []int{-1, arr[start]}
	}
	return []int{arr[end], arr[start]}
}

func dfsTransverse(node *TreeNode, arr *[]int) {
	if node.Left != nil {
		dfsTransverse(node.Left, arr)
	}
	*arr = append(*arr, node.Val)
	if node.Right != nil {
		dfsTransverse(node.Right, arr)
	}
}
