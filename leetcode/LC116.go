package leetcode

//   Definition for a Node1.
type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

func connect(root *Node1) *Node1 {
	if root == nil {
		return root
	}
	queue := []*Node1{root}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			current := queue[i]
			if i < n-1 {
				current.Next = queue[i+1]
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		queue = queue[n:]
	}
	return root
}
