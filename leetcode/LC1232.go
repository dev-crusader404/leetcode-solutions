package leetcode

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}
	dx := coordinates[1][0] - coordinates[0][0]
	dy := coordinates[1][1] - coordinates[0][1]

	for i := 2; i < len(coordinates); i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if dx*(y-coordinates[0][1]) != dy*(x-coordinates[0][0]) {
			return false
		}
	}
	return true
}
