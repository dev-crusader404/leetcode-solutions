package leetcode

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	fresh, rotten := analyse(grid)
	if fresh == 0 {
		return 0
	}
	if len(rotten) == 0 {
		return -1
	}
	return minTimeToRot(fresh, rotten, grid)
}

var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minTimeToRot(fresh int, queue, grid [][]int) int {
	minTime := 0

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			curr := queue[i]
			for d := 0; d < len(dir); d++ {
				row := dir[d][0] + curr[0]
				col := dir[d][1] + curr[1]
				if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
					continue
				}
				if grid[row][col] == 1 {
					grid[row][col] = 2
					fresh--
					queue = append(queue, []int{row, col})
				}
			}
		}
		queue = queue[n:]
		minTime++
		if fresh == 0 {
			return minTime
		}
	}
	return -1
}

func analyse(g [][]int) (int, [][]int) {
	var f int
	r := [][]int{}
	for i, v := range g {
		for j, a := range v {
			if a == 1 {
				f++
			}
			if a == 2 {
				r = append(r, []int{i, j})
			}
		}
	}
	return f, r
}

func RunLC994() {
	//a := [][]int{{2, 0, 1, 0, 0}, {1, 1, 0, 0, 2}, {0, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}
	a := [][]int{{1, 2, 1, 1, 2, 1, 1}}
	orangesRotting(a)
}
