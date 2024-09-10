package leetcode

import "container/heap"

type Items struct {
	Item string
	Freq int
}

type TopItem []*Items

func (l TopItem) Len() int      { return len(l) }
func (l TopItem) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l TopItem) Less(i, j int) bool {
	if l[i].Freq != l[j].Freq {
		return l[i].Freq > l[j].Freq
	}
	return l[i].Item < l[j].Item
}

func (l *TopItem) Push(x any) {
	d := x.(*Items)
	*l = append(*l, d)
}

func (l *TopItem) Pop() any {
	n := len(*l) - 1
	popped := (*l)[n]
	*l = (*l)[:n]
	return popped
}

func NewItems(i string, v int) *Items {
	return &Items{
		Item: i,
		Freq: v,
	}
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, s := range words {
		m[s]++
	}
	h := new(TopItem)
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, NewItems(k, v))
	}
	var res []string
	for k > 0 {
		x := heap.Pop(h).(*Items)
		res = append(res, x.Item)
		k--
	}
	return res
}
