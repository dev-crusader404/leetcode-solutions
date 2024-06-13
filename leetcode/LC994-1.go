package leetcode

import (
	"fmt"
	"math"
)

/*
	Given a 2D matrix containing, walls(-1), Gates(0) and EmptyCell(Inf).
	Find the minimum path to the gates from each cell.
*/

var Inf int = math.MaxInt

func WallAndGates(grid [][]int) [][]int {
	if len(grid) == 0 {
		return grid
	}
	start := findGates(grid)
	minPathToGates(start, &grid)
	return grid
}

// var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minPathToGates(stack [][]int, mat *[][]int) {
	for len(stack) > 0 {
		n := len(stack)
		curr := stack[n-1]
		stack = stack[:n-1]
		for i := 0; i < len(dir); i++ {
			row := curr[0] + dir[i][0]
			col := curr[1] + dir[i][1]
			now := (*mat)[curr[0]][curr[1]]
			if row < 0 || row >= len(*mat) || col < 0 || col >= len((*mat)[0]) ||
				(*mat)[row][col] == 0 || (*mat)[row][col] == -1 {
				continue
			}
			neighbor := (*mat)[row][col]
			if neighbor < (1 + now) {
				continue
			}
			(*mat)[row][col] = 1 + now
			stack = append(stack, []int{row, col})
		}
	}
}

func findGates(g [][]int) [][]int {
	var gateCords [][]int

	for i, v := range g {
		for j, w := range v {
			if w == 0 {
				gateCords = append(gateCords, []int{i, j})
			}
		}
	}
	return gateCords
}

func RunWallGate() {
	a := [][]int{{Inf, -1, 0, Inf}, {Inf, Inf, Inf, -1}, {Inf, -1, Inf, -1}, {0, -1, Inf, Inf}}
	fmt.Println(WallAndGates(a))
}
