package leetcode

type DLNode struct {
	Val   int
	Prev  *DLNode
	Next  *DLNode
	Child *DLNode
}

func flattenIterative(root *DLNode) *DLNode {
	if root == nil {
		return root
	}
	curr := root
	for curr != nil {
		if curr.Child != nil {
			tail := curr.Child
			for tail.Next != nil {
				tail = tail.Next
			}
			tail.Next = curr.Next
			if curr.Next != nil {
				curr.Next.Prev = tail
			}
			curr.Next = curr.Child
			curr.Child.Prev = curr
			curr.Child = nil
		}
		curr = curr.Next
	}
	return root
}

func flatten(root *DLNode) *DLNode {
	if root == nil {
		return root
	}
	curr := root

	flattenNode(curr)
	return root
}

func flattenNode(curr *DLNode) *DLNode {
	var end *DLNode
	for curr.Next != nil && curr.Child == nil {
		curr = curr.Next
	}
	if curr.Child != nil {
		end = flattenNode(curr.Child)
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
	for curr.Next != nil {
		curr = curr.Next
	}
	return curr
}
