// This is an implementation of priority queue for Graph data

package graph

type GraphPriorityQueue struct {
	vertex   *Vertex
	priority int
	index    int
}

type PriorityQueue []*GraphPriorityQueue

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	data := x.(*GraphPriorityQueue)
	data.index = n
	*pq = append(*pq, data)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	popped := old[n-1]
	old[n-1] = nil
	popped.index = -1
	*pq = old[:n-1]
	return popped
}

func NewGraphQueue(v *Vertex, p int) *GraphPriorityQueue {
	return &GraphPriorityQueue{
		vertex:   v,
		priority: p,
	}
}
