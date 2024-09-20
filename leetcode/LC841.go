package leetcode

func canVisitAllRooms(rooms [][]int) bool {
	isVis := make([]bool, len(rooms))
	stk := []int{0}
	isVis[0] = true
	for len(stk) > 0 {
		n := len(stk) - 1
		num := stk[n]
		stk = stk[:n]
		for _, j := range rooms[num] {
			if !isVis[j] {
				stk = append(stk, j)
				isVis[j] = true
			}
		}
	}
	for _, tv := range isVis {
		if !tv {
			return tv
		}
	}
	return true
}
