package leetcode

func smallestSubsequence(s string) string {
	lastIndex := make([]int, 26)
	seen := make([]bool, 26)
	for i, v := range s {
		lastIndex[v-'a'] = i
	}

	stk := make([]rune, 0)

	for i, c := range s {
		if seen[c-'a'] {
			continue
		}
		for len(stk) > 0 && c < stk[len(stk)-1] && lastIndex[stk[len(stk)-1]-'a'] > i {
			seen[stk[len(stk)-1]-'a'] = false
			stk = stk[:len(stk)-1]
		}
		seen[c-'a'] = true
		stk = append(stk, c)
	}
	return string(stk)
}
