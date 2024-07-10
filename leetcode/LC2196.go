package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
	if len(descriptions) == 0 || len(descriptions[0]) < 3 {
		return nil
	}
	childParent := make(map[int]struct{})
	m := make(map[int]*TreeNode)

	for _, val := range descriptions {
		childParent[val[1]] = struct{}{}
		var curr *TreeNode
		node, ok := m[val[0]]
		if !ok {
			curr = &TreeNode{Val: val[0]}
		} else {
			curr = node
		}

		if val[2] == 1 {
			if n, ok := m[val[1]]; ok {
				curr.Left = n
			} else {
				curr.Left = &TreeNode{Val: val[1]}
				m[val[1]] = curr.Left
			}
		} else {
			if n, ok := m[val[1]]; ok {
				curr.Right = n
			} else {
				curr.Right = &TreeNode{Val: val[1]}
				m[val[1]] = curr.Right
			}
		}
		m[val[0]] = curr
	}

	for _, v := range descriptions {
		if _, ok := childParent[v[0]]; !ok {
			return m[v[0]]
		}
	}
	return nil
}

func RunLC2196() {
	d := [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}
	createBinaryTree(d)
}
