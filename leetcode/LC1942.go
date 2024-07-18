package leetcode

import (
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
func (p PriQ) Less(i, j int) bool { return p[i].Seat < p[i].Seat }
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
}
