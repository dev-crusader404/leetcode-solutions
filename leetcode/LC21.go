package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var result *ListNode = &ListNode{}
	head := result
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			result.Next = list2
			list2 = list2.Next
		} else {
			result.Next = list1
			list1 = list1.Next
		}
		result = result.Next
	}

	if list1 != nil {
		result.Next = list1
	}
	if list2 != nil {
		result.Next = list2
	}
	return head.Next
}
