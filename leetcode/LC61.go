package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	current, runner := head, head

	for k > 0 {
		current = current.Next
		if current == nil {
			current = head
		}
		k--
	}

	if current == head {
		return current
	}
	for current.Next != nil {
		current = current.Next
		runner = runner.Next
	}

	newHead := runner.Next
	runner.Next = nil
	current.Next = head
	return newHead
}
