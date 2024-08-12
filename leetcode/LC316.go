package leetcode

func removeDuplicateLetters(s string) string {
	lastIndex := make([]int, 26)
	seen := make([]bool, 26)
	for i, v := range s {
		lastIndex[v-'a'] = i
	}

	stk := make([]rune, 0)

	for i, c := range s {
		for len(stk) > 0 && c < stk[len(stk)-1] && lastIndex[c-'a'] != i {
			seen[stk[len(stk)-1]-'a'] = false
			stk = stk[:len(stk)-1]
		}
		if seen[c-'a'] {
			continue
		}
		seen[c-'a'] = true
		stk = append(stk, c)
	}
	return string(stk)
}
