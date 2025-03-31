package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	if head == nil || len(nums) == 0 {
		return head
	}
	deleteMap := make(map[int]struct{})
	for _, n := range nums {
		deleteMap[n] = struct{}{}
	}
	var res, result *ListNode
	var c int
	for head != nil {
		if _, ok := deleteMap[head.Val]; !ok {
			if c == 0 {
				res = head
				result = res
			} else {
				res.Next = head
				res = res.Next
			}
			c++
		}
		head = head.Next
	}
	res.Next = nil
	return result
}

func RunLC3217() {
	l1 := &Node{}
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)
	l1.Add(4)
	l1.Add(5)
	l1.Add(6)
	x := []int{1, 2, 6}
	p := modifiedList(x, l1.head)
	p.PrintLinkedList()
}
