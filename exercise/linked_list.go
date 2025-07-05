package exercise

import "errors"

// Define List and dataNode types here.
// Note: The tests expect dataNode type to include an exported field with name Value to pass.

type List struct {
	head *dataNode
}

type dataNode struct {
	Value    interface{}
	next     *dataNode
	previous *dataNode
}

func NewList(elements ...interface{}) *List {
	l := &List{}
	for _, data := range elements {
		l.Push(data)
	}
	return l
}

func (n *dataNode) Next() *dataNode {
	return n.next
}

func (n *dataNode) Prev() *dataNode {
	return n.previous
}

func (l *List) Unshift(v interface{}) {
	curr := &dataNode{Value: v}
	if l == nil || l.head == nil {
		l.head = curr
		return
	}
	curr.next = l.head
	l.head.previous = curr
	l.head = curr
}

func (l *List) Push(v interface{}) {
	node := &dataNode{Value: v}
	if l.head == nil {
		l.head = node
	} else {
		tail := l.Last()
		tail.next = node
		node.previous = tail
	}
}

func (l *List) Shift() (interface{}, error) {
	if l == nil || l.head == nil {
		return nil, errors.New("empty list")
	}
	curr := l.head
	if curr.next != nil {
		curr.next.previous = nil
	}
	l.head = curr.next
	return curr.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l == nil || l.head == nil {
		return nil, errors.New("empty list")
	}
	tail := l.Last()
	if tail == l.head {
		popped := l.head.Value
		l.head = nil
		return popped, nil
	}
	tail.previous.next = nil
	tail.previous = nil
	return tail.Value, nil
}

func (l *List) Reverse() {
	if l.head == nil || l.Last() == l.First() {
		return
	}
	var prev *dataNode
	curr := l.head
	for curr != nil {
		runner := curr.next
		curr.next = prev
		curr.previous = runner
		prev = curr
		curr = runner
	}
	l.head = prev
}

func (l *List) First() *dataNode {
	return l.head
}

func (l *List) Last() *dataNode {
	curr := l.head
	for curr != nil && curr.next != nil {
		curr = curr.next
	}
	return curr
}
