package leetcode

import (
	"container/heap"
	"math"
)

var distance []int

type pq []int

func (a pq) Len() int {
	return len(a)
}

func (a pq) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a pq) Less(i, j int) bool {
	return distance[a[i]] < distance[a[j]]
}

func (a *pq) Push(x any) {
	*a = append(*a, x.(int))
}

func (a *pq) Pop() any {
	n := len(*a) - 1
	v := (*a)[n]
	(*a) = (*a)[:n]
	return v
}

func networkDelayTime2(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return 0
	}
	seen := make(map[int]struct{})
	distance = make([]int, n)
	for i := range distance {
		if i == k-1 {
			distance[i] = 0
		} else {
			distance[i] = math.MaxInt
		}
	}
	vertexMap := make(map[int][][2]int)
	for _, k := range times {
		u, v, w := k[0], k[1], k[2]
		vertexMap[u-1] = append(vertexMap[u-1], [2]int{v - 1, w})
	}
	h := &pq{}
	heap.Init(h)
	heap.Push(h, k-1)

	for h.Len() > 0 {
		curr := heap.Pop(h).(int)

		if _, ok := seen[curr]; ok {
			continue
		}
		seen[curr] = struct{}{}
		for _, vertex := range vertexMap[curr] {
			newDistance := vertex[1] + distance[curr]
			if newDistance < distance[vertex[0]] {
				distance[vertex[0]] = newDistance
				heap.Push(h, vertex[0])
			}
		}

	}
	var signalTime int
	for _, v := range distance {
		if v == math.MaxInt {
			return -1
		}
		signalTime = max(signalTime, v)
	}
	return signalTime
}
