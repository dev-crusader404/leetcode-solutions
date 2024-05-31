package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	current := head
	var newHead, revHead, lastTail *ListNode
	for current != nil {
		count := 0
		current = head
		for count < k && current != nil {
			current = current.Next
			count++
		}

		if count == k {
			revHead = reverseLinkedList(head, k)
			if newHead == nil {
				newHead = revHead
			}

			if lastTail != nil {
				lastTail.Next = revHead
			}
			lastTail = head
			head = current
		} else {
			lastTail.Next = head
		}
	}

	return newHead
}

func reverseLinkedList(ptr *ListNode, k int) *ListNode {
	var prev *ListNode
	for k > 0 {
		runner := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = runner
		k--
	}
	return prev
}

func (l *ListNode) PrintLinkedList() {
	if l == nil {
		return
	}
	for i := l; i != nil; i = i.Next {
		fmt.Printf("%d ", i.Val)
	}
	fmt.Println()
}

func RunLC25() {
	l := Node{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.head.PrintLinkedList()
	rev := reverseKGroup(l.head, 3)
	rev.PrintLinkedList()
}
