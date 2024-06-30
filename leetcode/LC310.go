package leetcode

import "math"

func findMinHeightTrees(n int, edges [][]int) []int {
	list := findAdjacency(n, edges)
	minH := math.MaxInt
	var result []int
	for i := 0; i < n; i++ {
		seen := make(map[int]struct{})
		h := findMinHeight(list, seen, math.MaxInt, i)
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

func findMinHeight(list [][]int, seen map[int]struct{}, height, node int) int {
	if _, ok := seen[node]; ok {
		return height
	}

	seen[node] = struct{}{}
	newHeight := height + 1
	for i := 0; i < len(list[node]); i++ {
		curr := list[node][i]
		height = min(height, findMinHeight(list, seen, newHeight, curr))
	}
	return height
}

func findAdjacency(n int, edges [][]int) [][]int {
	l := make([][]int, n)
	for _, v := range edges {
		l[v[0]] = append(l[v[0]], v[1])
		l[v[1]] = append(l[v[1]], v[0])
	}
	return l
}
