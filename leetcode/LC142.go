package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	var isCycle bool
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}
	for slow != head {
		slow = slow.Next
		head = head.Next
	}
	return head
}
