package leetcode

func removeStars(s string) string {
	stk := []rune{}
	for _, c := range s {
		if len(stk) > 0 && c == '*' {
			stk = stk[:len(stk)-1]
		} else if c != '*' {
			stk = append(stk, c)
		}
	}
	return string(stk)
}
