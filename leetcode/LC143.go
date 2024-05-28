package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev, current := slow, slow.Next

	// reversed the half part
	for current != nil {
		runner := current.Next
		current.Next = prev
		prev = current
		current = runner
	}

	for head != slow {
		temp := head.Next
		head.Next = prev
		head = head.Next
		prev = prev.Next
		head.Next = temp
	}

}
