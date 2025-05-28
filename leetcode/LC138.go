package leetcode

type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList(head *Node138) *Node138 {
	if head == nil {
		return head
	}
	nodeMap := make(map[*Node138]*Node138)
	curr := head

	for curr != nil {
		nodeMap[curr] = &Node138{Val: curr.Val}
		curr = curr.Next
	}
	curr = head
	for curr != nil {
		nodeMap[curr].Next = nodeMap[curr.Next]
		nodeMap[curr].Random = nodeMap[curr.Random]
		curr = curr.Next
	}
	return nodeMap[head]
}
