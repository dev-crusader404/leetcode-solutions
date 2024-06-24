package leetcode

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for k := 1; k < m; k++ {
		for l := 1; l < n; l++ {
			grid[k][l] += min(grid[k-1][l], grid[k][l-1])
		}
	}
	return grid[m-1][n-1]
}
