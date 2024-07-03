package leetcode

import (
	"fmt"
	"math"
)

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
	r := findMinHeightTrees2(6, a)
	fmt.Println(r)
}
