package leetcode

import (
	"container/heap"
	"fmt"
)

type maxHeap []float64

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(float64))
}

func (m *maxHeap) Pop() any {
	n := len(*m) - 1
	pop := (*m)[n]
	*m = (*m)[:n]
	return pop
}

func halveArray(nums []int) int {
	var sum float64
	var count int
	h := new(maxHeap)
	heap.Init(h)
	for _, v := range nums {
		sum += float64(v)
		heap.Push(h, float64(v))
	}
	half := sum / 2
	curr := sum
	for h.Len() > 0 && half > (sum-curr) {
		pop := heap.Pop(h).(float64)
		curr -= pop
		pop /= 2
		curr += pop
		heap.Push(h, pop)
		count++
	}
	return count
}

func RunLC2208() {
	a := []int{3, 8, 20}
	fmt.Println(halveArray(a))
}
