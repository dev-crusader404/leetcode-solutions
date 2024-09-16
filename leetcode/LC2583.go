package leetcode

import (
	"container/heap"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type qu []int

func (h qu) Len() int           { return len(h) }
func (h qu) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h qu) Less(i, j int) bool { return h[i] < h[j] }

func (h *qu) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *qu) Pop() any {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := []*TreeNode{root}
	minHeap := new(qu)
	heap.Init(minHeap)
	for len(queue) > 0 {
		n := len(queue)
		sum := 0
		for i := 0; i < n; i++ {
			node := queue[i]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if minHeap.Len() == k && (*minHeap)[0] < sum {
			heap.Pop(minHeap)
			heap.Push(minHeap, sum)
		}
		if minHeap.Len() < k {
			heap.Push(minHeap, sum)
		}
		queue = queue[n:]
	}
	if minHeap.Len() != k {
		return -1
	}
	return int64((*minHeap)[0])
}

func RunLC2583() {
	x := []any{5, 8, 9, 2, 1, 3, 7, 4, 6}
	t := BuildTree(x)
	fmt.Println(kthLargestLevelSum(t, 2))
}
