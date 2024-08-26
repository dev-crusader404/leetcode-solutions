package leetcode

import "fmt"

func equalPairs(grid [][]int) int {
	m := make(map[string]int)

	for _, row := range grid {
		var s string
		for _, col := range row {
			s += fmt.Sprintf("%d-", col)
		}
		if c, ok := m[s]; !ok {
			m[s] = 1
		} else {
			m[s] = c + 1
		}
	}
	var count int
	for i := range grid {
		var a string
		for j := range grid[0] {
			a += fmt.Sprintf("%d-", grid[j][i])
		}
		if c, ok := m[a]; ok {
			count += c
		}
	}
	return count
}
