package leetcode

import (
	"container/heap"
	"math"
)

type myNode struct {
	Node     int
	Distance int
}

func NewNode(x, y int) *myNode {
	return &myNode{
		Node:     x,
		Distance: y,
	}
}

type PriorityQ []*myNode

func (p PriorityQ) Len() int {
	return len(p)
}

func (p PriorityQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PriorityQ) Less(i, j int) bool {
	return p[i].Distance < p[j].Distance
}

func (p *PriorityQ) Push(x any) {
	d := x.(*myNode)
	*p = append(*p, d)
}

func (p *PriorityQ) Pop() any {
	n := len(*p) - 1
	pop := (*p)[n]
	*p = (*p)[:n]
	return pop
}

func networkDelayTime(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return 0
	}
	vertexMap := make(map[int][]myNode)
	for _, k := range times {
		u, v, w := k[0], k[1], k[2]
		vertexMap[u] = append(vertexMap[u], myNode{v, w})
	}

	seen := make(map[int]struct{})
	distance := make(map[int]int)
	initDistance(distance, n, k)
	queue := new(PriorityQ)
	heap.Init(queue)
	heap.Push(queue, NewNode(k, 0))
	for queue.Len() > 0 {
		curr := queue.Pop().(*myNode)

		if _, ok := seen[curr.Node]; ok {
			continue
		}
		seen[curr.Node] = struct{}{}
		for _, val := range vertexMap[curr.Node] {
			newDistance := distance[curr.Node] + val.Distance
			if distance[val.Node] > newDistance {
				distance[val.Node] = newDistance
				heap.Push(queue, NewNode(val.Node, distance[val.Node]))
			}
		}
	}

	if len(seen) != n {
		return -1
	}
	totalTime := math.MinInt
	for _, v := range distance {
		totalTime = max(totalTime, v)
	}
	return totalTime
}

func initDistance(distance map[int]int, n, k int) {
	for i := 1; i <= n; i++ {
		if i == k {
			distance[i] = 0
		} else {
			distance[i] = math.MaxInt
		}
	}
}
