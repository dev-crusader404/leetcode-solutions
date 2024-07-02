package leetcode

var chessMoves = [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}}

func knightProbability(n int, k int, row int, column int) float64 {
	if row < 0 || column < 0 || row >= n || column >= n {
		return float64(0)
	}
	if k == 0 {
		return float64(1)
	}
	var total float64
	for i := 0; i < len(chessMoves); i++ {
		dir := chessMoves[i]
		total += knightProbability(n, k-1, row+dir[0], column+dir[1]) / 8
	}

	return total
}
