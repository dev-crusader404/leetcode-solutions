package exercise

import "errors"

// Define the SimpleList and Element types here.
type SimpleList struct {
	head *Element
}

type Element struct {
	Id   int
	next *Element
}

func NewSimple(elements []int) *SimpleList {
	l := &SimpleList{}
	for _, ele := range elements {
		l.Push(ele)
	}
	return l
}

func (l *SimpleList) Size() int {
	if l == nil {
		return 0
	}
	size := 0
	curr := l.head
	for curr != nil {
		size++
		curr = curr.next
	}
	return size
}

func (l *SimpleList) Push(element int) {
	node := &Element{Id: element}
	if l.head == nil {
		l.head = node
	} else {
		curr := l.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node
	}
}

func (l *SimpleList) Pop() (int, error) {
	if l == nil || l.head == nil {
		return 0, errors.New("empty SimpleList")
	}
	curr := l.head
	if curr.next == nil {
		l.head = nil
		return curr.Id, nil
	}
	for curr.next.next != nil {
		curr = curr.next
	}
	popped := curr.next
	curr.next = nil
	return popped.Id, nil
}

func (l *SimpleList) Array() []int {
	var arr []int
	curr := l.head
	for curr != nil {
		arr = append(arr, curr.Id)
		curr = curr.next
	}
	return arr
}

func (l *SimpleList) Reverse() *SimpleList {
	var prev *Element
	if l.head == nil || l.head.next == nil {
		return l
	}
	curr := l.head
	for curr != nil {
		runner := curr.next
		curr.next = prev
		prev = curr
		curr = runner
	}
	l.head = prev
	return l
}
