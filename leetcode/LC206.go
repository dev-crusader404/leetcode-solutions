package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	var prev *ListNode
	for head != nil {
		current = head.Next
		head.Next = prev
		prev = head
		head = current
	}
	return prev
}
