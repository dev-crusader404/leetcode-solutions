package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	curr, count := head, 1
	var prev *ListNode
	for count < left && curr != nil {
		prev = curr
		curr = curr.Next
		count++
	}
	stop, c := curr, count
	for count <= right {
		stop = stop.Next
		count++
	}
	reversed := reversal(curr, stop, c, right)
	if prev != nil {
		prev.Next = reversed
	} else {
		return reversed
	}
	return head
}

func reversal(node, prev *ListNode, count, right int) *ListNode {
	for count <= right && node != nil {
		runner := node.Next
		node.Next = prev
		prev = node
		node = runner
		count++
	}
	return prev
}

func RunLC92() {
	l := Node{}
	for _, v := range []int{1, 2, 3, 4, 5} {
		l.Add(v)
	}
	reverseBetween(l.head, 3, 5)
}
