package leetcode

func minCostToMoveChips(position []int) int {
	var even int
	for _, v := range position {
		if v%2 == 0 {
			even++
		}
	}
	return min(even, len(position)-even)
}
