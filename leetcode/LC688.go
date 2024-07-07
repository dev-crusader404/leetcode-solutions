package leetcode

import "fmt"

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

func knightProbability2(n int, k int, row int, column int) float64 {
	var (
		dp                = make(map[string]float64)
		moves             = [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}}
		probabilitySolver func(k, row, column int) float64
	)

	probabilitySolver = func(k, row, column int) float64 {
		if row < 0 || column < 0 || row >= n || column >= n {
			return 0
		}
		if k == 0 {
			return 1
		}
		key := fmt.Sprintf("%d-%d-%d", k, row, column)
		if val, ok := dp[key]; ok {
			return val
		}
		var total float64
		for _, v := range moves {
			total += probabilitySolver(k-1, row+v[0], column+v[1]) * 0.125
		}
		dp[key] = total
		return dp[key]
	}
	return probabilitySolver(k, row, column)
}
