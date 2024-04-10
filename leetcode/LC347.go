package leetcode

import "container/heap"

type Item struct {
	num      int
	priority int
}

type PriorityQueue []*Item

func (p PriorityQueue) Len() int { return len(p) }
func (p PriorityQueue) Less(i, j int) bool {
	if p[i].priority == p[j].priority {
		return p[i].num > p[j].num
	} else {
		return p[i].priority > p[j].priority
	}
}
func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x any) {
	data := x.(*Item)
	*p = append(*p, data)
}
func (p *PriorityQueue) Pop() any {
	old := (*p)
	n := len(old)
	value := old[n-1]
	old[n-1] = nil
	(*p) = old[:n-1]
	return value
}

func NewItem(val, count int) *Item {
	return &Item{
		num:      val,
		priority: count,
	}
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) <= 0 {
		return []int{}
	}

	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}
	h := &PriorityQueue{}
	heap.Init(h)
	for i, v := range countMap {
		data := NewItem(i, v)
		heap.Push(h, data)
	}

	c := 0
	result := make([]int, 0)
	for c < k {
		v := heap.Pop(h).(*Item).num
		result = append(result, v)
		c++
	}
	return result
}
