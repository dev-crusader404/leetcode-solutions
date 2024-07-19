package leetcode

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	n := len(*m) - 1
	pop := (*m)[n]
	*m = (*m)[:n]
	return pop
}

func minOperations(nums []int, k int) int {
	var count int
	h := new(minHeap)
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}

	for h.Len() > 1 && (*h)[0] < k {
		n1, n2 := heap.Pop(h).(int), heap.Pop(h).(int)
		n1 = n1*2 + n2
		heap.Push(h, n1)
		count++
	}
	return count
}

func RunLC3066() {
	a := []int{2, 11, 10, 1, 3}
	fmt.Println(minOperations(a, 10))
}
