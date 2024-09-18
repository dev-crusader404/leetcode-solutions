package leetcode

import (
	"container/heap"
	"fmt"
)

type ClosestOrigin [][]int

func (c ClosestOrigin) Len() int      { return len(c) }
func (c ClosestOrigin) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c ClosestOrigin) Less(i, j int) bool {
	return (c[i][0]*c[i][0] + c[i][1]*c[i][1]) > (c[j][0]*c[j][0] + c[j][1]*c[j][1])
}

func (c *ClosestOrigin) Push(x any) {
	*c = append(*c, x.([]int))
}

func (c *ClosestOrigin) Pop() any {
	n := len(*c) - 1
	popped := (*c)[n]
	*c = (*c)[:n]
	return popped
}

func kClosest(points [][]int, k int) [][]int {
	if len(points) == 0 || k == 0 {
		return [][]int{}
	}
	h := new(ClosestOrigin)
	*h = append(*h, points[:k]...)
	heap.Init(h)
	for i := k; i < len(points); i++ {
		heap.Push(h, points[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return *h
}

func RunLC973() {
	p := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	fmt.Println(kClosest(p, 2))
}
