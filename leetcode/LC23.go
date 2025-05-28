package leetcode

import (
	"container/heap"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Priority []*ListNode

func (p Priority) Len() int           { return len(p) }
func (p Priority) Less(i, j int) bool { return p[i].Val < p[j].Val }
func (p Priority) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *Priority) Push(x any) {
	*p = append(*p, x.(*ListNode))
}

func (p *Priority) Pop() any {
	n := len(*p)
	old := (*p)[n-1]
	*p = (*p)[:n-1]
	return old
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	l := new(Priority)
	heap.Init(l)
	for _, node := range lists {
		if node != nil {
			heap.Push(l, node)
		}
	}
	var result *ListNode = &ListNode{}
	runner := result
	for l.Len() > 0 {
		node := heap.Pop(l).(*ListNode)
		runner.Next = node
		runner = runner.Next
		if node.Next != nil {
			heap.Push(l, node.Next)
		}
	}
	return result.Next
}

func RunLC23() {
	n := &Node{}
	n.Add(1)
	n.Add(4)
	n.Add(5)
	i := &Node{}
	i.Add(1)
	i.Add(3)
	i.Add(4)
	k := &Node{}
	k.Add(2)
	k.Add(6)
	var list []*ListNode
	list = append(list, n.head, i.head, k.head)
	r := mergeKLists(list)
	r.PrintLinkedList()
}
