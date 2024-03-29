package datastructure

import "fmt"

const Size = 10

type HashMap struct {
	ArrayMap [Size]*LinkListed
}

type LinkListed struct {
	head *ListNode
}

type ListNode struct {
	data string
	next *ListNode
}

// Insert function implements insertion of key to hashMap
func (h *HashMap) Insert(key string) {
	index := generateHash(key)
	h.ArrayMap[index].Insert(key)
}

// Insert function insert the key into Linkedlist
func (l *LinkListed) Insert(key string) {
	currentNode := &ListNode{data: key}
	currentNode.next = l.head
	l.head = currentNode
}

func initMap() *HashMap {
	hm := &HashMap{}
	for k := range hm.ArrayMap {
		hm.ArrayMap[k] = &LinkListed{}
	}
	return hm
}

func generateHash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % len(key)
}

func InitHashMap() {
	m := initMap()
	fmt.Println(m)

	l := &LinkListed{}
	l.Insert("Jack")
	l.Insert("Pete")
	l.Insert("Ryan")
}
