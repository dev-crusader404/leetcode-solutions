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
					if !result {
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
	l := []any{1, 10, 4, 3, nil, 7, 9, 12, 8, 6, nil, nil, 2}
	tree := BuildTree(l)
	fmt.Println(isEvenOddTree(tree))
}
