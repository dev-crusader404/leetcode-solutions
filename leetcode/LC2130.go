package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	if head == nil {
		return 0
	}
	slow := head
	fast := head
	var maximum int
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev *ListNode
	runner := head
	for head != slow {
		runner = head.Next
		head.Next = prev
		prev = head
		head = runner
	}
	for slow != nil {
		maximum = max(maximum, slow.Val+prev.Val)
		slow = slow.Next
		prev = prev.Next
	}
	return maximum
}
