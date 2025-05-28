package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	curr := head
	var res *ListNode = &ListNode{}
	node := res
	for curr != nil && curr.Next != nil {
		var flag bool
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
			flag = true
		}
		if !flag {
			res.Next = curr
			res = res.Next
		}
		curr = curr.Next
	}
	if curr != nil {
		res.Next = curr
		res = res.Next
	}
	if res.Next != nil {
		res.Next = nil
	}
	return node.Next
}
