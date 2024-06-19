package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{Val: -1}
	runner := newHead
	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			runner.Next = head
			runner = runner.Next
			head = head.Next
		}
	}
	runner.Next = head
	return newHead.Next
}
