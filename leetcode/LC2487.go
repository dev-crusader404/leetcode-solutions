package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var stack []*ListNode
	current := head

	for current != nil {
		for len(stack) > 0 && current.Val > stack[len(stack)-1].Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, current)
		current = current.Next
	}
	result := &ListNode{}
	resultCopy := result
	for i := 0; i < len(stack); i++ {
		resultCopy.Next = stack[i]
		resultCopy = resultCopy.Next
	}

	return result.Next
}

func removeNodes2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	result := &ListNode{}
	stack := make([]int, 0)
	var size int
	for head != nil {
		if size == 0 || stack[size-1] >= head.Val {
			stack = append(stack, head.Val)
			size++
		} else {
			for size > 0 && stack[size-1] < head.Val {
				size--
			}
			stack[size] = head.Val
			size++

			stack = stack[:size]
		}
		head = head.Next
	}
	resultCopy := result
	for i := 0; i < len(stack); i++ {
		resultCopy.Next = &ListNode{Val: stack[i]}
		resultCopy = resultCopy.Next
	}
	return result.Next
}

func RunLC2487() {
	l := &Node{}
	// for _, v := range []int{1, 1, 1, 1} {
	// 	l.Add(v)
	// }
	l.Add(5)
	l.Add(2)
	l.Add(13)
	l.Add(3)
	l.Add(8)

	fmt.Println(removeNodes(l.head))
}
