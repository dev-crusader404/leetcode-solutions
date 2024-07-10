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
	return searchTree(head, nil)
}

func getMidPoint(start, end *ListNode) *ListNode {
	fast, slow := start, start
	for fast != end && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
