package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(s)
	sort.Ints(g)
	var count int
	for i := 0; count < len(g) && i < len(s); i++ {
		if g[count] <= s[i] {
			count++
		}
	}
	return count
}
