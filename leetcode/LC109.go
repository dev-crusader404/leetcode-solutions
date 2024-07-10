package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	prev, slow, fast := head, head.Next, head.Next.Next

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	node := &TreeNode{Val: slow.Val}
	node.Right = sortedListToBST(slow.Next)
	node.Left = sortedListToBST(head)
	return node
}
