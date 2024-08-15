package leetcode

func scoreOfParentheses(s string) int {
	stk := []int{}
	score := 0
	for _, c := range s {
		if c == '(' {
			stk = append(stk, score)
			score = 0
		} else {
			score = stk[len(stk)-1] + max(2*score, 1)
			stk = stk[:len(stk)-1]
		}
	}
	return score
}
