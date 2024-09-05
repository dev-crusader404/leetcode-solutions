package leetcode

import "container/heap"

type SmallestInfiniteSet struct {
	index int
	arr   *smallPQ
}
type smallPQ []int

func (h smallPQ) Len() int           { return len(h) }
func (h smallPQ) Less(i, j int) bool { return h[i] < h[j] }
func (h smallPQ) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *smallPQ) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *smallPQ) Pop() any {
	n := len(*h) - 1
	ele := (*h)[n]
	*h = (*h)[:n]
	return ele
}

func (h *smallPQ) Search(x any) bool {
	n := x.(int)
	for i := range *h {
		if (*h)[i] == n {
			return true
		} else if (*h)[i] > n {
			break
		}
	}
	return false
}

func Constructor2336() SmallestInfiniteSet {
	h := new(smallPQ)
	heap.Init(h)
	return SmallestInfiniteSet{
		arr: h,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.arr.Len() > 0 {
		ele := heap.Pop(this.arr).(int)
		return ele
	}
	this.index++
	return this.index
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.index >= num && !this.arr.Search(num) {
		heap.Push(this.arr, num)
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
