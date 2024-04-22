package generics

import "fmt"

type genericList[T comparable] struct {
	data []T
}

func New[T comparable]() *genericList[T] {
	return &genericList[T]{
		data: make([]T, 0),
	}
}

func (l *genericList[T]) Insert(value T) {
	l.data = append(l.data, value)
}

func (l *genericList[T]) Get(index int) (V T) {
	if index > len(l.data)-1 {
		panic("input index out of range")
	}

	if index < 0 {
		panic(fmt.Sprintf("invalid index: %d", index))
	}
	return l.data[index]
}

func (l *genericList[T]) Remove(index int) (V T) {
	if index > len(l.data)-1 {
		panic("input index out of range")
	}

	if index < 0 {
		panic(fmt.Sprintf("invalid index: %d", index))
	}

	removedData := l.data[index]
	l.data = append(l.data[:index], l.data[index+1:]...)
	return removedData
}

func (l *genericList[T]) PrintList() {
	fmt.Println(l.data)
}

func GenerateList() {
	intList := New[int]()
	intList.Insert(12)
	intList.Insert(48)
	intList.Insert(78)
	intList.Insert(37)
	intList.Insert(10)
	intList.Insert(74)
	intList.Insert(55)
	intList.PrintList()
	fmt.Println(intList.Get(5))
	fmt.Println(intList.Get(3))
	fmt.Println(intList.Remove(2))
	fmt.Println(intList.Remove(5))
	intList.PrintList()
}
