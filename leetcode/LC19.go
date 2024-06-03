package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || (head.Next == nil && n == 1) {
		return nil
	}

	i, fast, slow := 0, head, head
	for i < n {
		fast = fast.Next
		i++
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil || (head.Next == nil && n == 1) {
		return nil
	}
	var count int
	current := head
	for current != nil {
		count++
		current = current.Next
	}

	ptr, index := head, 0
	var last *ListNode
	for ptr != nil {
		if (count - index) == n {
			if ptr == head {
				return head.Next
			} else {
				last.Next = ptr.Next
				break
			}
		}
		index++
		last = ptr
		ptr = ptr.Next
	}
	return head
}
