package leetcode

// small number after k digit removal
func removeKdigits(num string, k int) string {
	stk := []rune{}

	if len(num) == k {
		return "0"
	}

	for _, c := range num {
		for len(stk) > 0 && c < stk[len(stk)-1] && k > 0 {
			stk = stk[:len(stk)-1]
			k--
		}
		if len(stk) > 0 || c != '0' {
			stk = append(stk, c)
		}
	}

	for len(stk) > 0 && k > 0 {
		stk = stk[:len(stk)-1]
		k--
	}

	if len(stk) == 0 {
		return "0"
	}
	return string(stk)
}
