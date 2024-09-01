package leetcode

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stk := make([]int, 0)

	for i, temp := range temperatures {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] < temp {
			answer[stk[len(stk)-1]] = i - stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i)
	}
	return answer
}
