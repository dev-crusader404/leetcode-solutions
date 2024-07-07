package leetcode

import (
	"fmt"
	"math"
)

// Topological Sort T: O(E+N^2) S: O(N+E)
func findMinHeightTrees(n int, edges [][]int) []int {
	if n < 3 {
		var res []int
		for i := 0; i < n; i++ {
			res = append(res, i)
		}
		return res
	}

	list := make([][]int, n)
	indegree := make([]int, n)

	for _, v := range edges {
		indegree[v[0]]++
		indegree[v[1]]++
		list[v[0]] = append(list[v[0]], v[1])
		list[v[1]] = append(list[v[1]], v[0])
	}
	queue := make([]int, 0)
	for i, val := range indegree {
		if val == 1 {
			queue = append(queue, i)
		}
	}

	for n > 2 {
		size := len(queue)
		n -= size
		for i := 0; i < size; i++ {
			indegree[queue[i]]--
			for _, k := range list[queue[i]] {
				indegree[k]--
				if indegree[k] == 1 {
					queue = append(queue, k)
				}
			}
		}
		queue = queue[size:]
	}
	return queue
}

// This is recursive solution works. However, time complexity is greater so, might get time limit exceeded for large number of nodes
func findMinHeightTrees2(n int, edges [][]int) []int {
	var findMaxHeight func(height, node int) int
	list := make([][]int, n)
	for _, v := range edges {
		list[v[0]] = append(list[v[0]], v[1])
		list[v[1]] = append(list[v[1]], v[0])
	}
	minH := math.MaxInt
	var result []int
	for i := 0; i < n; i++ {
		seen := make(map[int]struct{})
		findMaxHeight = func(height, node int) int {
			seen[node] = struct{}{}
			newHeight := height + 1
			for i := 0; i < len(list[node]); i++ {
				curr := list[node][i]
				if _, ok := seen[curr]; !ok {
					height = max(height, findMaxHeight(newHeight, curr))
				}
			}
			return height
		}
		h := findMaxHeight(0, i)
		if h > minH {
			continue
		}
		if h < minH {
			result = result[:0]
		}
		result = append(result, i)
		minH = h
	}
	return result
}

func RunLC310() {
	// a := [][]int{{1, 0}, {1, 2}, {1, 3}}
	a := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	p := findMinHeightTrees(6, a)
	fmt.Println(p)
	r := findMinHeightTrees2(6, a)
	fmt.Println(r)
}
