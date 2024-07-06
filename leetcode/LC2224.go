package leetcode

import "time"

func convertTime(current string, correct string) int {
	from, _ := time.Parse("15:04", current)
	to, _ := time.Parse("15:04", correct)

	inMinutes := int(to.Sub(from).Minutes())

	var numOps int
	selectTime := []int{60, 15, 5, 1}

	for _, t := range selectTime {
		numOps += inMinutes / t
		inMinutes = inMinutes % t
		if inMinutes == 0 {
			break
		}
	}
	return numOps
}

func RunLC2224() {

}
