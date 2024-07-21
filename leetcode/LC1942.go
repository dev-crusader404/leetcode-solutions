package leetcode

import (
	"container/heap"
	"fmt"
	"sort"
)

type SeatAvailability struct {
	Seat          int
	AvailableTime int
}

func NewSeatAvailability(s, t int) *SeatAvailability {
	return &SeatAvailability{Seat: s, AvailableTime: t}
}

type PriQ []*SeatAvailability

func (p PriQ) Len() int           { return len(p) }
func (p PriQ) Less(i, j int) bool { return p[i].AvailableTime < p[j].AvailableTime }
func (p PriQ) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *PriQ) Push(x any) {
	*p = append(*p, x.(*SeatAvailability))
}

func (p *PriQ) Pop() any {
	n := len(*p) - 1
	pop := (*p)[n]
	*p = (*p)[:n]
	return pop
}

func (p *PriQ) Peek() any {
	return (*p)[0]
}

type IntHeaps []int

func (h IntHeaps) Len() int           { return len(h) }
func (h IntHeaps) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeaps) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeaps) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeaps) Pop() any {
	n := len(*h) - 1
	pop := (*h)[n]
	*h = (*h)[:n]
	return pop
}

func smallestChair(times [][]int, targetFriend int) int {
	start, end := times[targetFriend][0], times[targetFriend][1]
	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})

	seat := &IntHeaps{}
	heap.Init(seat)
	for i := 0; i < len(times); i++ {
		heap.Push(seat, i)
	}

	k := new(PriQ)
	heap.Init(k)
	for _, v := range times {
		for k.Len() > 0 && k.Peek().(*SeatAvailability).AvailableTime <= v[0] {
			heap.Push(seat, heap.Pop(k).(*SeatAvailability).Seat)
		}
		if v[0] == start && v[1] == end {
			break
		} else {
			heap.Push(k, NewSeatAvailability(heap.Pop(seat).(int), v[1]))
		}
	}
	return heap.Pop(seat).(int)
}

func RunLC1942() {
	// t := [][]int{{1, 4}, {2, 3}, {4, 6}}
	t := [][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}}
	fmt.Println(smallestChair(t, 6))
}
