package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	current := slow

	// reversed the half part
	for current != nil {
		runner := current.Next
		current.Next = prev
		prev = current
		current = runner
	}

	for head.Next != slow {
		temp := head.Next
		head.Next = prev
		head = head.Next
		prev = prev.Next
		head.Next = temp
		head = head.Next
	}

	for prev != nil {
		head.Next = prev
		prev = prev.Next
		head = head.Next
	}
}

func RunLC143() {
	l1 := Node{}
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)
	l1.Add(4)
	l1.Add(5)
	reorderList(l1.head)
}
