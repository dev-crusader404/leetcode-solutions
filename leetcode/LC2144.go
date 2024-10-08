package leetcode

import "sort"

func minimumCost(cost []int) int {
	var price int
	sort.Ints(cost)
	c := 0
	for i := len(cost) - 1; i >= 0; i-- {
		c++
		if c%3 == 0 {
			continue
		}
		price += cost[i]
	}
	return price
}
