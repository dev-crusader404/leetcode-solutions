package leetcode

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var count int
	for i, x := range grid {
		for j, y := range x {
			if y == '1' {
				count++
				countIsland(grid, i, j)
			}
		}
	}
	return count
}

func countIsland(grid [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	countIsland(grid, row+1, col)
	countIsland(grid, row-1, col)
	countIsland(grid, row, col+1)
	countIsland(grid, row, col-1)
}
