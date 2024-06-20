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
	var prev, dumm *ListNode
	for count < left && curr != nil {
		prev = curr
		curr = curr.Next
		count++
	}
	end := curr
	for count <= right && curr != nil {
		runner := curr.Next
		curr.Next = dumm
		dumm = curr
		curr = runner
		count++
	}
	end.Next = curr
	if prev != nil {
		prev.Next = dumm
	} else {
		return dumm
	}
	return head
}

func RunLC92() {
	l := Node{}
	for _, v := range []int{1, 2, 3, 4, 5} {
		l.Add(v)
	}
	reverseBetween(l.head, 3, 5)
}
