package leetcode

import (
	"container/list"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	stk1 := list.New()
	stk2 := list.New()
	result := list.New()

	for l1 != nil {
		stk1.PushFront(l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stk2.PushFront(l2.Val)
		l2 = l2.Next
	}

	res := &ListNode{}
	dummy := res

	carry := 0
	s1 := stk1.Front()
	s2 := stk2.Front()
	for s1 != nil || s2 != nil || carry != 0 {
		sum := 0
		if s1 != nil {
			sum += s1.Value.(int)
		}

		if s2 != nil {
			sum += s2.Value.(int)
		}

		sum += carry
		carry = sum / 10
		sum %= 10
		result.PushFront(sum)
		if s1 != nil {
			s1 = s1.Next()
		}
		if s2 != nil {
			s2 = s2.Next()
		}
	}

	for e := result.Front(); e != nil; e = e.Next() {
		dummy.Next = &ListNode{Val: e.Value.(int)}
		dummy = dummy.Next
	}
	return res.Next
}

func RunLC445() {
	l1 := &Node{}
	l1.Add(7)
	l1.Add(2)
	l1.Add(4)
	l1.Add(3)
	l2 := &Node{}
	l2.Add(5)
	l2.Add(6)
	l2.Add(4)

	fmt.Println(addTwoNumbers2(l1.head, l2.head))
}
