package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect2(root *Node1) *Node1 {
	if root == nil {
		return root
	}
	queue := []*Node1{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			currNode := queue[i]
			if i < n-1 {
				currNode.Next = queue[i+1]
			}
			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}

			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}
		queue = queue[n:]
	}
	return root
}
