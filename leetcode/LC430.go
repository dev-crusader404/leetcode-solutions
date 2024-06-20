package leetcode

type DLNode struct {
	Val   int
	Prev  *DLNode
	Next  *DLNode
	Child *DLNode
}

func flatten(root *DLNode) *DLNode {
	if root == nil {
		return root
	}
	curr := root

	flattenNode(curr, nil)
	return root
}

func flattenNode(curr, parent *DLNode) *DLNode {
	var end *DLNode
	for curr.Next != nil && curr.Child == nil {
		curr = curr.Next
	}
	if curr.Child != nil {
		end = flattenNode(curr.Child, curr)
		end.Next = curr.Next
		curr.Child.Prev = curr
	}
	if curr.Next != nil {
		curr.Next.Prev = end
	}
	if curr.Child != nil {
		curr.Next = curr.Child
	}
	curr.Child = nil
	if parent == nil {
		return curr
	}
	for curr.Next != nil {
		curr = curr.Next
	}
	return curr
}
