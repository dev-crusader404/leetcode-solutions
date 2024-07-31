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
func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	level := 0
	queue := []*TreeNode{root}
	var result bool = true
	for len(queue) > 0 {
		n := len(queue)
		isEven := (level%2 == 0)
		prev := 0
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if (node.Val%2 == 1) == isEven {
				result = true
				if i > 0 {
					result = isEven == (prev < node.Val)
					if prev == node.Val || !result {
						return false
					}
				}
			} else {
				return false
			}
			prev = node.Val
		}
		level++
		queue = queue[n:]
	}
	return result
}

func RunLC1609() {
	l := []any{13, 34, 32, 23, 25, 27, 29, 44, 40, 36, 34, 30, 30, 28, 26, 3, 7, 9, 11, 15, 17, 21, 25, nil, nil, 27, 31, 35, nil, 37, nil, 30, nil, 26, nil, nil, nil, 24, nil, 20, 16, 12, 10, nil, nil, 8, nil, nil, nil, nil, nil, 6, nil, nil, nil, nil, nil, 15, 19, nil, nil, nil, nil, 23, nil, 27, 29, 33, 37, nil, nil, nil, nil, nil, nil, 48, nil, nil, nil, 46, nil, nil, nil, 42, 38, 34, 32, nil, nil, nil, nil, 19}
	tree := BuildTree(l)
	fmt.Println(isEvenOddTree(tree))
}
