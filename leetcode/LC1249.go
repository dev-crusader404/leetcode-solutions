package leetcode

import (
	"strings"
)

// Using byte array
func minRemoveToMakeValid2(s string) string {
	c := []byte(s)
	var stk []int
	var x byte
	for i, v := range c {
		if v == '(' {
			stk = append(stk, i)
		} else if v == ')' && len(stk) > 0 {
			n := len(stk) - 1
			stk = stk[:n]
		} else if v == ')' {
			c[i] = x
		}
	}

	for _, v := range stk {
		c[v] = x
	}

	// var st bytes.Buffer
	// for _, v := range c {
	// 	if v != x {
	// 		st.WriteByte(v)
	// 	}
	// }
	// return st.String()
	return strings.ReplaceAll(string(c), "\u0000", "")
}

func minRemoveToMakeValid(s string) string {
	var openCount, closeCount int
	removal := make(map[int]struct{})
	for i, v := range s {
		if v == '(' {
			openCount++
		} else if v == ')' {
			if openCount > 0 {
				openCount--
			} else {
				removal[i] = struct{}{}
			}
		}
	}

	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == ')' {
			closeCount++
		} else if s[j] == '(' {
			if closeCount > 0 {
				closeCount--
			} else {
				removal[j] = struct{}{}
			}
		}
	}
	var sb strings.Builder
	for i := range s {
		if _, ok := removal[i]; !ok {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
