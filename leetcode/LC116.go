package leetcode

//   Definition for a Node.
type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

func connect(root *Node1) *Node1 {
	if root != nil {
		return root
	}
	list := []*Node1{root}

	for len(list) > 0 {
		n := len(list)
		for i := 0; i < n; i++ {
			current := list[i]
			if i != n-1 {
				current.Next = list[i+1]
			}
			if current.Left != nil {
				list = append(list, current.Left)
			}
			if current.Right != nil {
				list = append(list, current.Right)
			}
		}
		list = list[n:]
	}
	return root
}
