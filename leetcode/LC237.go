package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode237(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
