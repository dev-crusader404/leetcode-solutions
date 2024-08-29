package leetcode

import "container/heap"

type minheap []int

func (m minheap) Len() int           { return len(m) }
func (m minheap) Swap(a, b int)      { m[a], m[b] = m[b], m[a] }
func (m minheap) Less(a, b int) bool { return m[a] < m[b] }

func (m *minheap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minheap) Pop() any {
	n := len(*m) - 1
	x := (*m)[n]
	*m = (*m)[:n]
	return x
}

func maximumProduct(nums []int, k int) int {
	var mod int = 1e9 + 7
	h := new(minheap)
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
	}
	product := 1
	for k > 0 {
		ele := heap.Pop(h).(int)
		heap.Push(h, ele+1)
		k--
	}
	for h.Len() > 0 {
		product = product * heap.Pop(h).(int) % mod
	}
	return product
}
