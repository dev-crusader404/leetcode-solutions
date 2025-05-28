package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}
	return head
}
