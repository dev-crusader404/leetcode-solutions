package leetcode

import "strings"

func removeOuterParentheses(s string) string {
	c := 0
	var sb strings.Builder
	for _, v := range s {
		if v == '(' {
			if c > 0 {
				sb.WriteRune(v)
			}
			c++
		}

		if v == ')' {
			if c > 1 {
				sb.WriteRune(v)
			}
			c--
		}
	}
	return sb.String()
}
