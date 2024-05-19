package leetcode

import "fmt"

type Node struct {
	head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *Node) Add(x int) {
	if l.head == nil {
		l.head = &ListNode{Val: x}
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &ListNode{Val: x}
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum, carry := 0, 0

	result := &ListNode{}
	actual := result

	for l1 != nil || l2 != nil || carry != 0 {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum = (x + y + carry)
		carry = sum / 10
		sum %= 10
		result.Next = &ListNode{Val: sum}
		result = result.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return actual.Next
}

func RunLC2() {
	l1 := &Node{}
	l1.Add(2)
	l1.Add(4)
	l1.Add(3)
	l2 := &Node{}
	l2.Add(5)
	l2.Add(6)
	l2.Add(4)

	fmt.Println(addTwoNumbers(l1.head, l2.head))
}
