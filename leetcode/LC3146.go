package leetcode

import "math"

func findPermutationDifference(s string, t string) int {
	m := make(map[rune]int)
	for i, v := range s {
		m[v] = i
	}
	var diff int

	for i, k := range t {
		c := m[k]
		diff += int(math.Abs(float64(c - i)))
	}
	return diff
}
