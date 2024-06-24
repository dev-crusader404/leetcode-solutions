package leetcode

import "strings"

func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string)
	for _, v := range knowledge {
		m[v[0]] = v[1]
	}
	var left int
	var mapBound bool
	var bul strings.Builder

	for right, val := range s {
		if val == '(' {
			left = right
			mapBound = true
		} else if val == ')' {
			if k, ok := m[s[left+1:right]]; ok {
				bul.WriteString(k)
			} else {
				bul.WriteRune('?')
			}
			mapBound = false
			continue
		}
		if !mapBound {
			bul.WriteRune(val)
		}
	}
	return bul.String()
}
