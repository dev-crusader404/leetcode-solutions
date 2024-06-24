package leetcode

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	(*h) = append((*h), x.(int))
}

func (h *IntHeap) Pop() any {
	n := len(*h) - 1
	pop := (*h)[n]
	(*h) = (*h)[:n]
	return pop
}

func findKthLargest2(nums []int, k int) int {
	h := IntHeap(nums[:k])
	heap.Init(&h)

	for _, v := range nums[k:] {
		if h[0] < v {
			heap.Pop(&h)
			heap.Push(&h, v)
		}
	}
	return h[0]
}
